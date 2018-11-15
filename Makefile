.ONESHELL:

# Basic go commands
GOCMD=go
ZIPCMD=zip
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=Handle

all: build compress

compress:
	cd bin; $(ZIPCMD) -r $(BINARY_NAME).zip . -x *.git*

build:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -v -ldflags "-d -s -w" -a -tags netgo -installsuffix net -o bin/$(BINARY_NAME) src/main.go
