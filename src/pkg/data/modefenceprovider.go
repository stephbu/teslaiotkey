package data

type MockFenceProvider struct {
	location LatLong
	distance float64
}

func NewMockFenceProvider(location LatLong, distance float64) *MockFenceProvider {
	return &MockFenceProvider{location: location, distance: distance}
}

func (mock *MockFenceProvider) GetDistance() float64 {
	return mock.distance
}

func (mock *MockFenceProvider) GetLocation() LatLong {
	return mock.location
}
