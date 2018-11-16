# Tesla IOT Key

## Overview
Goal of this project is to create some code to support 
using an Amazon IOT Button as an unlock my Tesla only
when it is parked in my garage.

## Problem Statement
When I leave for work, I need to unplug my Tesla,
however this cannot be done until the car is awoke and unlocked e.g.
by opening the door, or using the app to unlock and wake the car, 
then walking back to unplug the car.

Further still I only want this button to function while the 
car is in the garage or near the house.  I do not want the 
car to unlock when it is not at home.

## Solution Overview
Use an Amazon IoT button to will send a message 
to unlock the Tesla via the Tesla API.

The code inside of this repository contains a basic implementation
of a Lambda function to respond to Amazon IOT Button, and its related 
IoTButtonEvent.

## References
[Tesla API Reference](https://tesla-api.timdorr.com)

[Geographic Coordinate System](https://en.wikipedia.org/wiki/Geographic_coordinate_system)

[Amazon IOT Button Gen.2](https://www.amazon.com/AWS-IoT-Button-2nd-Generation/dp/B01KW6YCIM)

[AWS Lambda Golang Reference](https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model.html)

[AWS Lambda Deployment Package in Go](https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html)

[IotButtonEvent Sample](https://muzigram.muzigen.net/2018/06/iotbutton-go-firebase-golang-lambda.html)

## Notes
### Getting Started with Golang with Lambda

#### Magic names for Golang binaries
The recipe to get started with Golang and Lambda is superficially a little messy and took some trial and error to get right.
First the error messages are misleading - ```no such file or directory```.

Turns out that Lambda expects the binary name, and the function name 
for your lambda entrypoint to be the same.  30mins of my
life I'll never get back.

#### Golang IotButton Samples
Moreover there are no official samples for IotButtonEvent written in Go at this point that appear in Google.
If they do exist, they are buried.

#### Security and Lambda
I'll pass credentials for my Tesla API key into the lambda through environment variables. Investigate how to securely attach those to the Lambda definition.
Amazon Key Management API enables creation of secrets, and a client-side API to decode secrets within the code.  Cipher key is exchanged externally as part of the
IAM role/context of the Lambda


#### Tesla API Libraries
Looking around for Tesla API libraries and instructions for the Tesla Owners REST API.
Started with a fork of [ElectricGopher](https://github.com/stephbu/electricgopher/),
