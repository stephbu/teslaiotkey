package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stephbu/teslaiotkey/src/pkg/data"
	"github.com/stephbu/teslaiotkey/src/pkg/logging"
)

var config *data.Configuration
var teslaInstance *data.TeslaCarProvider
var homeFence *data.GeofenceProvider

func Handle(ctx context.Context, button events.IoTButtonEvent) (string, error) {

	var err error
	config, err = data.LoadConfigFromEnv()
	if err != nil {
		return "", err
	}

	teslaInstance = data.NewTeslaCarProvider(config)
	homeFence = data.NewGeofenceProvider(config)

	handlerContext := logging.CreateRequestContext(ctx)
	logger := logging.WithContext(handlerContext)

	logger.
		WithField("ClickType", button.ClickType).
		WithField("SerialNumber", button.SerialNumber).
		WithField("Voltage", button.BatteryVoltage).
		Info("Request received")

	command := config.ClickMap[button.ClickType]
	if command == "" {
		return "", errors.New(fmt.Sprintf("No commands mapped to '%s'", button.ClickType))
	} else {
		logger.
			WithField("ClickType", button.ClickType).
			WithField("Command", command)
	}

	insideFence, err := homeFence.IsInFence(ctx, teslaInstance)
	if err != nil {
		return "", err
	}

	if insideFence {
		logger.Info("inside of fence")

		err = teslaInstance.Invoke(handlerContext, command)
		if err != nil {
			return "", err
		}
	} else {
		logger.Info("outside of fence")
		// log outside fence
	}

	logger.
		WithField("ClickType", button.ClickType).
		WithField("SerialNumber", button.SerialNumber).
		WithField("VIN", teslaInstance.VIN).
		Info("Request Complete")

	return fmt.Sprintf("Unlocked VIN:%v", teslaInstance.VIN), nil
}
