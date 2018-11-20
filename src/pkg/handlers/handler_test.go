package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stephbu/teslaiotkey/src/pkg/logging"
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
