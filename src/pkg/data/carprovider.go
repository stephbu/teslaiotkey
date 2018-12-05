package data

type CarProvider interface {
	LocationProvider
	LockProvider
	CommandProvider
}
