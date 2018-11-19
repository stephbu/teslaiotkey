package data

import (
	"errors"
	"os"
	"strconv"
)

const (
	TESLA_VIN      = "TESLA_VIN"
	TESLA_USERNAME = "TESLA_USERNAME"
	TESLA_PASSWORD = "TESLA_PASSWORD"
	FENCE_LATLONG  = "FENCE_LATLONG"
	FENCE_RADIUS   = "FENCE_RADIUS"
)

type Configuration struct {
	VIN          string
	Username     string
	Password     string
	FenceLatLong *LatLong
	FenceRadius  float64 // Radius of the geofence in meters
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
	result.FenceRadius, err = strconv.ParseFloat(fenceRadius, 64)
	if err != nil {
		return nil, errors.New("invalid fence environment variable")
	}

	return
}
