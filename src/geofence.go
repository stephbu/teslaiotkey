package main

import (
	"github.com/kellydunn/golang-geo"
	"github.com/stephbu/teslaiotkey/src/pkg/data"
)

func getCarToFenceDistanceMeters(fence data.FenceProvider, car data.LocationProvider) float64 {
	fenceLocation := fence.GetLocation()
	carLocation := car.GetLocation()

	return geo.NewPoint(fenceLocation.Lat, fenceLocation.Long).GreatCircleDistance(geo.NewPoint(carLocation.Lat, carLocation.Long)) * 1000
}
