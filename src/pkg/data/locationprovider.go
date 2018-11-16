package data

type LocationProvider interface {
	GetLocation() LatLong
}

type FenceProvider interface {
	LocationProvider
	GetDistance() float64
}
