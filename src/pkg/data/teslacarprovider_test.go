package data

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewTeslaCarProviderInitialize(t *testing.T) {
	config, err := LoadConfigFromEnv()
	assert.Nil(t, err)

	teslaCarProvider := NewTeslaCarProvider(config)
	err = teslaCarProvider.initialize()

	if err != nil {
		t.Error(err)
	}
}

func TestLoadConfigFromEnv1(t *testing.T) {

	const VIN = "vin"
	const USERNAME = "username"
	const PASSWORD = "password"
	const LATLONG = "47.642744,-122.112782"
	const RADIUS = "30"

	os.Setenv(TESLA_VIN, VIN)
	os.Setenv(TESLA_USERNAME, USERNAME)
	os.Setenv(TESLA_PASSWORD, PASSWORD)
	os.Setenv(FENCE_LATLONG, LATLONG)
	os.Setenv(FENCE_RADIUS, RADIUS)

	config, err := LoadConfigFromEnv()
	assert.Nil(t, err)
	assert.Equal(t, config.VIN, VIN)
	assert.Equal(t, config.Username, USERNAME)
	assert.Equal(t, config.Password, PASSWORD)
	assert.Equal(t, config.FenceLatLong.Lat, float64(47.642744))
	assert.Equal(t, config.FenceLatLong.Long, float64(-122.112782))
	assert.Equal(t, config.FenceRadius, float64(30))
}

func TestLoadConfigFromEnv2(t *testing.T) {

	const VIN = ""
	const USERNAME = "username"
	const PASSWORD = "password"
	const LATLONG = "47.642744,-122.112782"
	const RADIUS = "30"

	os.Setenv(TESLA_VIN, VIN)
	os.Setenv(TESLA_USERNAME, USERNAME)
	os.Setenv(TESLA_PASSWORD, PASSWORD)
	os.Setenv(FENCE_LATLONG, LATLONG)
	os.Setenv(FENCE_RADIUS, RADIUS)

	config, err := LoadConfigFromEnv()
	assert.Error(t, err)
	assert.Nil(t, config)
}

func TestLoadConfigFromEnv3(t *testing.T) {

	const VIN = "vin"
	const USERNAME = ""
	const PASSWORD = "password"
	const LATLONG = "47.642744,-122.112782"
	const RADIUS = "30"

	os.Setenv(TESLA_VIN, VIN)
	os.Setenv(TESLA_USERNAME, USERNAME)
	os.Setenv(TESLA_PASSWORD, PASSWORD)
	os.Setenv(FENCE_RADIUS, RADIUS)
	os.Setenv(FENCE_LATLONG, LATLONG)

	config, err := LoadConfigFromEnv()
	assert.Error(t, err)
	assert.Nil(t, config)
}
