# Installation
## Prerequisites
There are a number of prerequisite items required to get started.

1) Amazon Web Services Account - sign up at [AWS Console](https://aws.amazon.com).  
- <b>YES</b> use the same account as your Amazon.com shopping account.
- <b>DO</b> enable multi-factor authentication such as the [Authy app](https://authy.com) to protect your account and your car
- <b>FREE TIER</b> should be sufficient for your development and day-to-day use
2) Purchase an Amazon IOT Button
- [AWS IoT Button](https://aws.amazon.com/iotbutton/)
3) Install Go Language
4) Install deps
4) Clone this repo
- ```$GOPATH/src/github.com/stephbu/teslaiotkey``` 
## Building Binaries
Project has a makefile in the root of the repo.  ```$REPO/make``` currently I've only validated the Makefile on OS/X
- <b>HELP:</b> welcome PRs to make build and CI work for other OS's/Platforms
- Expects completely packaged Unix binaries in zip form.
- binaries dropped as ```$REPO/bin/Handle```
- zipped binaries dropped as ```$REPO/bin/Handle.zip```

## Configure IOT Device
Follow steps in packaging.
1) Register Device using Amazon IOT phone app.
2) Setup button Wifi network connection.
3) Test your device at [Amazon IOT Portal](https://us-west-2.console.aws.amazon.com/iotv2)

## Create Lambda Function
<i><b>HELP</b> [Open issue](https://github.com/stephbu/teslaiotkey/issues/3) to write scripts to install and configure the Lambda function,
create test event etc.</i>
1) Create new function in [Lambda console](https://us-west-2.console.aws.amazon.com/lambda/)
2) Add IOT Trigger
3) Upload zip function code 
    - from Building Binaries step
    - Go Runtime 1.x
    - function name ```Handle```
4) Add environment variables (no quotes, no leading/trailing spaces)
    ```
    TESLA_VIN=13 character Vehicle VIN Number   e.g. 5YJ3E1EB3JFxxxxxx
    TESLA_USERNAME=your tesla account username  e.g. test@gmail.com
    TESLA_PASSWORD=your tesla password          e.g. Password1!
    FENCE_LATLONG=homelat,homelong              e.g. 47.629272,-122.147589 
    FENCE_RADIUS=activationDistance(meters)     e.g. 30
    ```
5) Create test event

    ```
    {
      "SerialNumber": "G012345678901234",
      "ClickType": "SINGLE",
      "BatteryVoltage": "1592mV"
    }
    ```
6) Save everything.
7) Hit Test - your car should unlock.