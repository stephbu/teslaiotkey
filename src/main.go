package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stephbu/teslaiotkey/src/pkg/data"
)

var config *data.Configuration
var teslaInstance *data.TeslaCarProvider
var homeFence *data.HomeGeofenceProvider

func Handle(button events.IoTButtonEvent) (string, error) {

	distance, err := getCarToFenceDistanceMeters(homeFence, teslaInstance)
	if err != nil {
		return "", err
	}

	if distance < homeFence.GetDistance() {
		teslaInstance.SetState(data.UNLOCKED)
	}

	return fmt.Sprintf("Go hello from %s!", button.SerialNumber), nil
}

func main() {

	config, err := data.LoadConfigFromEnv()
	if err != nil {
		return
	}

	teslaInstance = data.NewTeslaCarProvider(config)
	homeFence = data.NewHomeGeofenceProvider(config)

	lambda.Start(Handle)
}
