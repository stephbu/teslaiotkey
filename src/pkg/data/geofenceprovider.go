package data

type GeofenceProvider struct {
	distance float64
	location *LatLong
}

func NewGeofenceProvider(config *Configuration) *GeofenceProvider {
	result := &GeofenceProvider{config.FenceRadius, config.FenceLatLong}
	return result
}

func (fence *GeofenceProvider) GetDistance() float64 {
	return fence.distance
}

func (fence *GeofenceProvider) GetLocation() (LatLong, error) {
	return *fence.location, nil
}
