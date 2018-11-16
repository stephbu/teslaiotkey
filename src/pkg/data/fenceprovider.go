package data

type FenceProvider interface {
	LocationProvider
	GetDistance() float64
}
