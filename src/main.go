package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stephbu/teslaiotkey/src/pkg/handlers"
)

func main() {
	lambda.Start(handlers.Handle)
}
