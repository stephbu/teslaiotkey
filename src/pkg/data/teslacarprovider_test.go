package data

import (
	"github.com/pkg/errors"
	"os"
	"testing"
)

const (
	TESLA_VIN      = "TESLA_VIN"
	TESLA_USERNAME = "TESLA_USERNAME"
	TESLA_PASSWORD = "TESLA_PASSWORD"
)

type Configuration struct {
	VIN      string
	Username string
	Password string
}

func LoadConfigFromEnv() (result *Configuration, err error) {

	result = &Configuration{}
	result.VIN = os.Getenv(TESLA_VIN)
	result.Username = os.Getenv(TESLA_USERNAME)
	result.Password = os.Getenv(TESLA_PASSWORD)

	// check parameters supplied
	if result.VIN == "" || result.Username == "" || result.Password == "" {
		return &Configuration{}, errors.New("missing environment variables")
	}

	return
}

func TestNewTeslaCarProviderInitialize(t *testing.T) {

	config, err := LoadConfigFromEnv()
	if err != nil {
		t.Error(err)
	}

	teslaCarProvider := NewTeslaCarProvider(config.VIN, config.Username, config.Password)
	err = teslaCarProvider.initialize()

	if err != nil {
		t.Error(err)
	}
}
