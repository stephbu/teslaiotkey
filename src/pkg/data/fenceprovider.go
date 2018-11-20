package data

type FenceProvider interface {
	LocationProvider
	GetDistance() float64
	IsInFence(pointLocation LocationProvider) (bool, error)
}
