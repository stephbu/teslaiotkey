package data

type MockCarProvider struct {
	location LatLong
}

func NewMockCarProvider(location LatLong) *MockCarProvider {
	return &MockCarProvider{location: location}
}

func (mock *MockCarProvider) GetLocation() LatLong {
	return mock.location
}
