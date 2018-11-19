package data

type HomeGeofenceProvider struct {
	distance float64
	location *LatLong
}

func NewHomeGeofenceProvider(config *Configuration) *HomeGeofenceProvider {
	result := &HomeGeofenceProvider{config.FenceRadius, config.FenceLatLong}
	return result
}

func (fence *HomeGeofenceProvider) GetDistance() float64 {
	return fence.distance
}

func (fence *HomeGeofenceProvider) GetLocation() (*LatLong, error) {
	return fence.location, nil
}
