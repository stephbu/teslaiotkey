package data

import "github.com/kellydunn/golang-geo"

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

func (fence *GeofenceProvider) IsInFence(pointLocation LocationProvider) (bool, error) {

	distance, err := FenceToPointDistance(fence, pointLocation)
	if err != nil {
		return false, err
	}

	if distance > fence.distance {
		return false, nil
	}

	return true, nil

}

func FenceToPointDistance(fenceProvider FenceProvider, pointProvider LocationProvider) (float64, error) {

	fence, err := fenceProvider.GetLocation()
	if err != nil {
		return 0, err
	}

	point, err := pointProvider.GetLocation()
	if err != nil {
		return 0, err
	}

	distance := geo.NewPoint(fence.Lat, point.Long).GreatCircleDistance(geo.NewPoint(point.Lat, point.Long)) * 1000
	return distance, nil
}
