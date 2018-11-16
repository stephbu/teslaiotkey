package data

type LocationProvider interface {
	GetLocation() (LatLong, error)
}
