package data

import "github.com/pkg/errors"

type MockFenceProvider struct {
	location    LatLong
	distance    float64
	throwErrors bool
}

func NewMockFenceProvider(location LatLong, distance float64, throwErr bool) *MockFenceProvider {
	return &MockFenceProvider{location: location, distance: distance, throwErrors: throwErr}
}

func (mock *MockFenceProvider) GetDistance() float64 {
	return mock.distance
}

func (mock *MockFenceProvider) GetLocation() (LatLong, error) {
	if mock.throwErrors {
		return LatLong{}, errors.New("mock error during GetLocation")
	}
	return mock.location, nil
}

func (mock *MockFenceProvider) IsInFence(pointLocation LocationProvider) (bool, error) {

	distance, err := FenceToPointDistance(mock, pointLocation)

	if err != nil {
		return false, err
	}
	if distance > mock.distance {
		return false, nil
	}
	return true, nil

}
