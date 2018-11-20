package handlers

import (
	"context"
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

	insideFence, err := homeFence.IsInFence(ctx, teslaInstance)
	if err != nil {
		return "", err
	}

	logger := logging.WithContext(handlerContext)

	if insideFence {
		logger.Info("inside of fence")
		err = teslaInstance.Unlock(handlerContext)
		if err != nil {
			return "", err
		}
	} else {
		logger.Info("outside of fence")
		// log outside fence
	}

	logger.
		WithField("SerialNumber", button.SerialNumber).
		WithField("VIN", teslaInstance.VIN).
		Info("Request Complete")

	return fmt.Sprintf("Unlocked VIN:%v", teslaInstance.VIN), nil
}
