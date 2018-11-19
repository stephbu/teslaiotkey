package data

type LockProvider interface {
	SetState(state LockState) (LockState, error)
}

type LockState int

const ( // iota is reset to 0
	UNKNOWN  LockState = iota
	LOCKED   LockState = iota
	UNLOCKED LockState = iota
)
