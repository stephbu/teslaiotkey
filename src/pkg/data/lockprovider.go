package data

import "context"

type LockProvider interface {
	Unlock(ctx context.Context) error
}
