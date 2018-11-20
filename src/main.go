package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stephbu/teslaiotkey/src/pkg/handlers"
)

func main() {
	lambda.Start(Handle)
}

func Handle(ctx context.Context, button events.IoTButtonEvent) (string, error) {
	return handlers.Handle(ctx, button)
}
