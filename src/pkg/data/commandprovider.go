package data

import "context"

type CommandProvider interface {
	Invoke(ctx context.Context, command string) error
}
