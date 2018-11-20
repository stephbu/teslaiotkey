package data

import (
	"context"
	"errors"
)

type MockCarProvider struct {
	location    LatLong
	throwErrors bool
}

func NewMockCarProvider(location LatLong, throwErr bool) *MockCarProvider {
	return &MockCarProvider{location: location, throwErrors: throwErr}
}

func (mock *MockCarProvider) GetLocation(ctx context.Context) (LatLong, error) {
	if mock.throwErrors {
		return LatLong{}, errors.New("mock error during GetLocation")
	}
	return mock.location, nil
}

func (mock *MockCarProvider) Unlock() error {
	if mock.throwErrors {
		return errors.New("mock error during Unlock")
	}
	return nil
}
