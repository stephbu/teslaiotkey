package data

import (
	"context"
	"github.com/kellydunn/golang-geo"
	"github.com/stephbu/teslaiotkey/src/pkg/logging"
)

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

func (fence *GeofenceProvider) GetLocation(ctx context.Context) (LatLong, error) {
	return *fence.location, nil
}

func (fence *GeofenceProvider) IsInFence(ctx context.Context, pointLocation LocationProvider) (bool, error) {

	distance, err := FenceToPointDistance(ctx, fence, pointLocation)
	if err != nil {
		return false, err
	}

	if distance > fence.distance {
		return false, nil
	}

	return true, nil

}

func FenceToPointDistance(ctx context.Context, fenceProvider FenceProvider, pointProvider LocationProvider) (float64, error) {

	fence, err := fenceProvider.GetLocation(ctx)
	if err != nil {
		return 0, err
	}

	point, err := pointProvider.GetLocation(ctx)
	if err != nil {
		return 0, err
	}

	distance := geo.NewPoint(fence.Lat, point.Long).GreatCircleDistance(geo.NewPoint(point.Lat, point.Long)) * 1000
	logging.WithContext(ctx).Printf("FenceToPointDistance=%v meters", distance)
	return distance, nil
}
