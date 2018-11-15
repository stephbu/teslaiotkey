package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handle(button events.IoTButtonEvent) (string, error) {
	//ctx := context.Background()
	return fmt.Sprintf("Go hello from %s!", button.SerialNumber ), nil
}

func main() {
	lambda.Start(Handle)
}
