package data

type LockProvider interface {
	GetState() (LockState, error)
	SetState(state LockState) (LockState, error)
}

type LockState int

const ( // iota is reset to 0
	UNKNOWN  LockState = iota
	LOCKED   LockState = iota
	UNLOCKED LockState = iota
)
