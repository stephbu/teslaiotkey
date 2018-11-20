package data

import "context"

type FenceProvider interface {
	LocationProvider
	GetDistance() float64
	IsInFence(ctx context.Context, pointLocation LocationProvider) (bool, error)
}
