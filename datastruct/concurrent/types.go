package concurrent

import "context"

type Queue[T any] interface {
	Enqueue(ctx context.Context, data any) error

	Dequeue(ctx context.Context) (any, error)

	IsEmpty() bool

	Size() uint64

	Len() uint64

	IsFull() bool
}
