package data

import "context"

type LocationProvider interface {
	GetLocation(ctx context.Context) (LatLong, error)
}
