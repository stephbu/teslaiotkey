package data

import "errors"

type MockCarProvider struct {
	location    LatLong
	throwErrors bool
}

func NewMockCarProvider(location LatLong, throwErr bool) *MockCarProvider {
	return &MockCarProvider{location: location, throwErrors: throwErr}
}

func (mock *MockCarProvider) GetLocation() (LatLong, error) {
	if mock.throwErrors {
		return LatLong{}, errors.New("mock error during GetLocation")
	}
	return mock.location, nil
}

func (mock *MockCarProvider) SetState(state LockState) (LockState, error) {
	if mock.throwErrors {
		return UNKNOWN, errors.New("mock error during SetState")
	}
	return UNKNOWN, nil
}
