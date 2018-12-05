package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stephbu/teslaiotkey/src/pkg/data"
	"github.com/stephbu/teslaiotkey/src/pkg/logging"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandle(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	button := events.IoTButtonEvent{SerialNumber: "foo", BatteryVoltage: "bar", ClickType: "SINGLE"}
	ctx := logging.CreateRequestContext(context.Background())

	Handle(ctx, button)
}

func TestHandle1(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	const VIN = "vin"
	const USERNAME = "username"
	const PASSWORD = "password"
	const LATLONG = "47.642744,-122.112782"
	const RADIUS = "30"

	os.Setenv(data.TESLA_VIN, VIN)
	os.Setenv(data.TESLA_USERNAME, USERNAME)
	os.Setenv(data.TESLA_PASSWORD, PASSWORD)
	os.Setenv(data.FENCE_LATLONG, LATLONG)
	os.Setenv(data.FENCE_RADIUS, RADIUS)

	button := events.IoTButtonEvent{SerialNumber: "foo", BatteryVoltage: "bar", ClickType: "UNKNOWNTYPE"}
	ctx := logging.CreateRequestContext(context.Background())

	_, err := Handle(ctx, button)
	assert.Error(t, err)
}

func TestHandle2(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	const VIN = "vin"
	const USERNAME = "username"
	const PASSWORD = "password"
	const LATLONG = "47.642744,-122.112782"
	const RADIUS = "30"
	const CLICKMAP = ",,"

	os.Setenv(data.TESLA_VIN, VIN)
	os.Setenv(data.TESLA_USERNAME, USERNAME)
	os.Setenv(data.TESLA_PASSWORD, PASSWORD)
	os.Setenv(data.FENCE_LATLONG, LATLONG)
	os.Setenv(data.FENCE_RADIUS, RADIUS)
	os.Setenv(data.CLICK_MAP, CLICKMAP)

	button := events.IoTButtonEvent{SerialNumber: "foo", BatteryVoltage: "bar", ClickType: data.CLICKTYPE_SINGLE}
	ctx := logging.CreateRequestContext(context.Background())

	_, err := Handle(ctx, button)
	assert.Error(t, err)
}

func TestHandle3(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	const VIN = "vin"
	const USERNAME = "username"
	const PASSWORD = "password"
	const LATLONG = "47.642744,-122.112782"
	const RADIUS = "30"
	const CLICKMAP = "unlock,,"

	os.Setenv(data.TESLA_VIN, VIN)
	os.Setenv(data.TESLA_USERNAME, USERNAME)
	os.Setenv(data.TESLA_PASSWORD, PASSWORD)
	os.Setenv(data.FENCE_LATLONG, LATLONG)
	os.Setenv(data.FENCE_RADIUS, RADIUS)
	os.Setenv(data.CLICK_MAP, CLICKMAP)

	button := events.IoTButtonEvent{SerialNumber: "foo", BatteryVoltage: "bar", ClickType: data.CLICKTYPE_SINGLE}
	ctx := logging.CreateRequestContext(context.Background())

	_, err := Handle(ctx, button)
	assert.Error(t, err)
}
