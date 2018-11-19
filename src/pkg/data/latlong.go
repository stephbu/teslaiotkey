package data

import (
	"errors"
	"strconv"
	"strings"
)

type LatLong struct {
	Lat  float64
	Long float64
}

// take strings formatted as [lat float],[long float]
func LatLongFromString(latLongInput string) (*LatLong, error) {

	// split by commas
	components := strings.Split(latLongInput, ",")

	// assert
	if len(components) != 2 {
		return nil, errors.New("invalid string format")
	}

	// try conversion

	if components[0] == "" {
		return nil, errors.New("invalid Latitude")
	}
	lat, err := strconv.ParseFloat(components[0], 64)
	if err != nil {
		return nil, errors.New("invalid Latitude, conversion failed")
	}

	if components[1] == "" {
		return nil, errors.New("invalid Longitude")
	}

	long, err := strconv.ParseFloat(components[1], 64)
	if err != nil {
		return nil, errors.New("invalid Longitude, conversion failed")
	}

	return &LatLong{Lat: lat, Long: long}, nil
}
