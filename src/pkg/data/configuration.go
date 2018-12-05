package data

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	TESLA_VIN            = "TESLA_VIN"
	TESLA_USERNAME       = "TESLA_USERNAME"
	TESLA_PASSWORD       = "TESLA_PASSWORD"
	TESLA_WAKEUP_TIMEOUT = "TESLA_WAKEUP_TIMEOUT"
	FENCE_LATLONG        = "FENCE_LATLONG"
	FENCE_RADIUS         = "FENCE_RADIUS"
	CLICK_MAP            = "CLICK_MAP" // Comma separated functions for single, double, long clicks

	WakeupTimeoutSecondsDefault = int(20)
	FenceRadiusMetersDefault    = float64(10)

	CLICKTYPE_SINGLE = "SINGLE"
	CLICKTYPE_DOUBLE = "DOUBLE"
	CLICKTYPE_LONG   = "LONG"
)

type ClickMap map[string]string

var ClickMapDefault = ClickMap{
	CLICKTYPE_SINGLE: "unlock",
	CLICKTYPE_DOUBLE: "unlock",
	CLICKTYPE_LONG:   "unlock",
}

type Configuration struct {
	VIN           string
	Username      string
	Password      string
	FenceLatLong  *LatLong
	FenceRadius   float64  // Radius of the geofence in meters
	WakeupTimeout int      // Wakeup Timeout Seconds
	ClickMap      ClickMap // Map of function to button press type
}

func LoadConfigFromEnv() (result *Configuration, err error) {

	result = &Configuration{}
	result.VIN = os.Getenv(TESLA_VIN)
	result.Username = os.Getenv(TESLA_USERNAME)
	result.Password = os.Getenv(TESLA_PASSWORD)

	// check car parameters supplied
	if result.VIN == "" || result.Username == "" || result.Password == "" {
		return nil, errors.New("missing car environment variables")
	}

	fenceLatLong := os.Getenv(FENCE_LATLONG)
	result.FenceLatLong, err = LatLongFromString(fenceLatLong)
	if err != nil {
		return nil, errors.New("invalid fence environment variable")
	}

	fenceRadius := os.Getenv(FENCE_RADIUS)
	if len(fenceRadius) > 0 {
		result.FenceRadius, err = strconv.ParseFloat(fenceRadius, 64)
		if err != nil {
			return nil, errors.New("invalid fence environment variable")
		}
	} else {
		result.FenceRadius = FenceRadiusMetersDefault
	}

	wakeupTimeout := os.Getenv(TESLA_WAKEUP_TIMEOUT)
	if len(wakeupTimeout) > 0 {
		result.WakeupTimeout, err = strconv.Atoi(wakeupTimeout)
		if err != nil {
			return nil, errors.New("invalid wakeup timeout environment variable")
		}
	} else {
		result.WakeupTimeout = WakeupTimeoutSecondsDefault
	}

	clickMap := os.Getenv(CLICK_MAP)
	if len(clickMap) == 0 {
		result.ClickMap = ClickMapDefault
	} else {
		values := strings.Split(clickMap, ",")
		result.ClickMap = ClickMap{}

		if len(values) >= 1 {
			result.ClickMap[CLICKTYPE_SINGLE] = values[0]
		}
		if len(values) >= 2 {
			result.ClickMap[CLICKTYPE_DOUBLE] = values[1]
		}
		if len(values) >= 3 {
			result.ClickMap[CLICKTYPE_LONG] = values[2]
		}
	}

	return
}
