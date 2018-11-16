package data

import "github.com/stephbu/electricgopher/api"

type TeslaCarProvider struct {
	// External parameters
	VIN         string // API interactions needs to be tied
	Credentials UsernamePasswordCredential

	// Internal state
	initialized   bool
	vehicleId     int64
	hasDriveState bool
	apiEndpoint   string

	client *api.Client
}

const (
	TESLA_CLIENT_ID     = "81527cff06843c8634fdc09e8ac0abefb46ac849f38fe1e431c2ef2106796384"
	TESLA_CLIENT_SECRET = "c7257eb71a564034f9419ee651c7d0e5f7aa6bfbd18bafb5c5c033b093bb2fa3"
)

func NewTeslaCarProvider(vin string, username string, password string) *TeslaCarProvider {
	teslaCarProvider := &TeslaCarProvider{VIN: vin, Credentials: UsernamePasswordCredential{Username: username, Password: password}}
	return teslaCarProvider
}

func (tesla *TeslaCarProvider) GetLocation() (LatLong, error) {
	return LatLong{}, nil
}

func (tesla *TeslaCarProvider) GetState() (LockState, error) {
	tesla.initialize()
	// TODO: Implement read lock state
	return UNKNOWN, nil
}

func (tesla *TeslaCarProvider) SetState(state LockState) (LockState, error) {
	tesla.initialize()
	// TODO: Implement write lock state
	return UNKNOWN, nil
}

// Lazy initialization function.
func (tesla *TeslaCarProvider) initialize() error {
	if tesla.initialized {
		return nil
	}

	tesla.client = api.NewClient(TESLA_CLIENT_ID, TESLA_CLIENT_SECRET, tesla.Credentials.Username, tesla.Credentials.Password, "https://owner-api.teslamotors.com/", nil)
	vehicles, err := tesla.client.GetVehicles()
	if err != nil {
		return err
	}

	// Get Vehicle
	for _, vehicle := range vehicles.Response {
		if vehicle.Vin == tesla.VIN {
			tesla.vehicleId = vehicle.Id
		}
	}

	tesla.initialized = true
	return nil
}

type UsernamePasswordCredential struct {
	Username string
	Password string
}
