# Tesla IOT Key

## Overview
Goal of this project is to create some code to support 
using a smart-button to unlock a Tesla only parked nearby.

## Status & Installation
Code is functional, but still work in progress (see [Open Issues](https://github.com/stephbu/teslaiotkey/issues))
Follow instructions here for [Installation](https://github.com/stephbu/teslaiotkey/blob/master/INSTALLATION.md)

## Problem Statement
When I leave for work, I need to unplug my Tesla,
however this cannot be done until the car is awoke and unlocked e.g.
by opening the door, or using the app to unlock and wake the car, 
then walking back to unplug the car.

Further still I only want this button to function while the 
car is in the garage or near the house.  I do not want the 
car to unlock when it is not at home.

## Requirements
- Enable a smart-button to unlock the car
- Prevent the car from being unlocked when out of eye-sight from my house or garage.
- Be robust enough to function every day without fail
- Be cheap enough to use everyday

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

#### Tesla OAuth Secrets
Tesla's API Gateway has a set of secrets to enable access.  Cadence of change as yet unknown. They were initially extracted from the 
Android application approximately 1yr ago  (notes from Reddit here: https://www.reddit.com/r/teslamotors/comments/72ilu3/a_few_tidbits_from_disassembling_the_tesla/)
Working secrets appear to be here - https://pastebin.com/pS7Z6yyP

#### Tesla API Libraries
Looking around for Tesla API libraries and instructions for the Tesla Owners REST API. [Tim Dorr's Tesla Ruby API](https://github.com/timdorr/tesla-api) is conceptually complete
and has a great cassette reference of the interactions with the Tesla API server.  I used Tim's code in conjunction with [Nathan Beyer's](https://nbeyer.io) partial Golang port [ElectricGopher](https://github.com/nbeyer/electricgopher) to 
port to get started on a functional client - currently working in ```vehicle-fix``` branch - [https://github.com/stephbu/electricgopher/]()

Some things I've fixed so far: 
- Socket leakage 
- HTTP POST client support
- Vehicle Command Library
     - Wakeup
     - Unlock


