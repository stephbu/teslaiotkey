package data

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stephbu/electricgopher/api"
	"github.com/stephbu/teslaiotkey/src/pkg/logging"
	"strconv"
	"time"
)

type TeslaCarProvider struct {
	// External parameters
	VIN         string // API interactions are tied from VIN to an internal Vehicle ID number
	Credentials UsernamePasswordCredential

	// Internal state
	initialized     bool    // indicates that the client is correctly initialized
	vehicleId       string  // internal Tesla vehicle identifier
	vehicleState    string  // last phone home state after wakeup
	vehicleLocation LatLong // last vehicle location after wakeup
	hasDriveState   bool
	apiEndpoint     string // API prefix

	client *api.Client
}

const (
	// TODO: move somewhat static tesla secrets into environment variables, just in case.
	TESLA_CLIENT_ID       = "81527cff06843c8634fdc09e8ac0abefb46ac849f38fe1e431c2ef2106796384"
	TESLA_CLIENT_SECRET   = "c7257eb71a564034f9419ee651c7d0e5f7aa6bfbd18bafb5c5c033b093bb2fa3"
	VEHICLE_STATE_ASLEEP  = "asleep"
	VEHICLE_STATE_ONLINE  = "online"
	VEHICLE_STATE_OFFLINE = "offline"
)

func NewTeslaCarProvider(config *Configuration) *TeslaCarProvider {
	teslaCarProvider := &TeslaCarProvider{VIN: config.VIN, Credentials: UsernamePasswordCredential{Username: config.Username, Password: config.Password}}

	return teslaCarProvider
}

func (tesla *TeslaCarProvider) GetLocation(ctx context.Context) (LatLong, error) {

	err := tesla.initialize(ctx)
	if err != nil {
		return LatLong{}, err
	}

	return tesla.vehicleLocation, nil
}

func (tesla *TeslaCarProvider) Unlock(ctx context.Context) error {

	err := tesla.initialize(ctx)
	if err != nil {
		return err
	}

	output, err := tesla.client.Unlock(tesla.vehicleId)
	if err != nil {
		return err
	}

	logging.WithContext(ctx).Printf("TeslaCarProvider::Unlock() response %+v", output)
	if !output.Response.Result {
		return errors.New(fmt.Sprintf("Unlock failed - %v", output.Response.Reason))
	}

	return nil
}

// Lazy initialization function.
func (tesla *TeslaCarProvider) initialize(ctx context.Context) error {

	logger := logging.WithContext(ctx)

	if tesla.initialized {
		return nil
	}

	tesla.client = api.NewClient(TESLA_CLIENT_ID, TESLA_CLIENT_SECRET, tesla.Credentials.Username, tesla.Credentials.Password, "https://owner-api.teslamotors.com", logger)
	vehicles, err := tesla.client.GetVehicles()
	if err != nil {
		logger.Error(err)
		return err
	}

	// Get Vehicle matching the VIN
	matchingCar := false
	for _, vehicle := range vehicles.Response {
		if vehicle.Vin == tesla.VIN {
			tesla.vehicleId = strconv.FormatInt(vehicle.Id, 10)
			tesla.vehicleState = vehicle.State
			logger.Printf("VIN:%v Vehicle ID: %+v", tesla.VIN, tesla.vehicleId)
			matchingCar = true
		}
	}

	if !matchingCar {
		return errors.New(fmt.Sprintf("no car matching VIN:%+v", tesla.VIN))
	}

	awake := false
	retryCount := 15
	retryThrottle := time.Tick(time.Second)

	for !awake && retryCount > 0 {

		logger.Printf("Current Car State: %+v (remaining connection attempts: %d)", tesla.vehicleState, retryCount)

		switch tesla.vehicleState {
		case VEHICLE_STATE_OFFLINE, VEHICLE_STATE_ASLEEP:
			logger.Print("Waking car")
			wakeUpResponse, err := tesla.client.WakeUp(tesla.vehicleId)
			if err != nil {
				logger.Printf("WakeUp %+v", err)
				return err
			}
			tesla.vehicleState = wakeUpResponse.Response.State
			logger.Printf("Wake Up Response %+v", wakeUpResponse)
		case VEHICLE_STATE_ONLINE:
			awake = true
		}

		if !awake {
			retryCount--
			<-retryThrottle
		}
	}

	if !awake {
		return errors.New("Could not wake car in allotted time")
	}

	drive, err := tesla.client.GetDriveState(tesla.vehicleId)
	if err != nil {
		logger.Printf("GetDriveState %+v", err)
		return err
	}
	logger.Printf("%+v", drive.Response)
	tesla.vehicleLocation = LatLong{Lat: drive.Response.Latitude, Long: drive.Response.Longitude}

	vehicle, err := tesla.client.GetVehicleState(tesla.vehicleId)
	if err != nil {
		logger.Printf("GetVehicleState %+v", err)
		return err
	}
	logger.Printf("%+v", vehicle.Response)

	tesla.initialized = true
	return nil
}

type UsernamePasswordCredential struct {
	Username string
	Password string
}
