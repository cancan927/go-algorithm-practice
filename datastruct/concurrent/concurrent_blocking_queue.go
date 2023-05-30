package concurrent

import (
	"context"
	"sync"
)

type BlockingQueue[T any] struct {
	lock     *sync.Mutex
	data     []T
	notFull  chan struct{}
	notEmpty chan struct{}
}

func NewBlockingQueue[T any](maxSize int) *BlockingQueue[T] {
	m := &sync.Mutex{}
	return &BlockingQueue[T]{
		data:     make([]T, 0, maxSize),
		lock:     m,
		notFull:  make(chan struct{}),
		notEmpty: make(chan struct{}),
	}
}

func (q *BlockingQueue[T]) Enqueue(ctx context.Context, data any) error {
	// 上来先判断是否超时
	if ctx.Err() != nil {
		return ctx.Err()
	}
	// 加锁
	q.lock.Lock()

	for q.IsFull() {
		q.lock.Unlock()
		select {
		case <-q.notFull:
			q.lock.Lock()
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	q.data = append(q.data, data)
	q.notEmpty <- struct{}{}

	return nil
}

func (q *BlockingQueue[T]) Dequeue(ctx context.Context) (any, error) {
	// 判断超时
	if ctx.Err() != nil {
		var t T
		return t, ctx.Err()
	}
	q.lock.Lock()
	for q.IsEmpty() {
		q.lock.Unlock()
		select {
		case <-q.notEmpty:
			q.lock.Lock()
		case <-ctx.Done():
			var t T
			return t, ctx.Err()
		}

	}

	t := q.data[0]
	q.data = q.data[1:]
	q.notFull <- struct{}{}
	return t, nil
}

func (q *BlockingQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *BlockingQueue[T]) Size() uint64 {
	return uint64(len(q.data))
}

func (q *BlockingQueue[T]) Len() uint64 {
	return 0
}

func (q *BlockingQueue[T]) IsFull() bool {
	return true
}
