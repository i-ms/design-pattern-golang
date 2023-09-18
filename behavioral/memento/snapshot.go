package memento

type Snapshot[T any] struct {
	State T
}
