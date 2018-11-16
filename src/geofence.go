package main

import (
	"github.com/kellydunn/golang-geo"
	"github.com/stephbu/teslaiotkey/src/pkg/data"
)

func getCarToFenceDistanceMeters(fence data.FenceProvider, car data.CarProvider) (float64, error) {
	fenceLocation, err := fence.GetLocation()
	if err != nil {
		return 0, nil
	}

	carLocation, err := car.GetLocation()
	if err != nil {
		return 0, nil
	}

	return geo.NewPoint(fenceLocation.Lat, fenceLocation.Long).GreatCircleDistance(geo.NewPoint(carLocation.Lat, carLocation.Long)) * 1000, nil
}
