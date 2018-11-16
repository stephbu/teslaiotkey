package data

import (
	"os"
	"testing"
)

const (
	TESLA_VIN      = "TESLA_VIN"
	TESLA_USERNAME = "TESLA_USERNAME"
	TESLA_PASSWORD = "TESLA_PASSWORD"
)

func TestNewTeslaCarProvider(t *testing.T) {

	teslaVin := os.Getenv(TESLA_VIN)
	teslaUsername := os.Getenv(TESLA_USERNAME)
	teslaPassword := os.Getenv(TESLA_PASSWORD)

	teslaCarProvider := NewTeslaCarProvider(teslaVin, teslaUsername, teslaPassword)
	teslaCarProvider.initialize()
}
