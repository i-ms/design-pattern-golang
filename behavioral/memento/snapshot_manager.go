package memento

import "sync"

type SnapshotManager[T any] struct {
	snapshots []Snapshot[T]
	version   int
	lock      sync.Mutex
}

func NewSnapshotManager[T any]() *SnapshotManager[T] {
	return &SnapshotManager[T]{
		snapshots: make([]Snapshot[T], 0),
		version:   -1,
	}
}

// AddSnapshot adds a snapshot to the snapshot manager.
func (sm *SnapshotManager[T]) AddSnapshot(s Snapshot[T]) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	// Check if there are snapshots that are undone.
	// i.e snapshots that are reverted to previous state
	if sm.version < len(sm.snapshots)-1 {
		sm.snapshots = sm.snapshots[:sm.version+1]
	}

	sm.snapshots = append(sm.snapshots, s)
	sm.version++
}

// GetPreviousSnapshot returns the previous snapshot from the snapshot manager.
func (sm *SnapshotManager[T]) GetPreviousSnapshot() Snapshot[T] {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if sm.version > 0 {
		sm.version--
	}

	return sm.snapshots[sm.version]
}

// GetNextSnapshot returns the next snapshot from the snapshot manager.
func (sm *SnapshotManager[T]) GetNextSnapshot() Snapshot[T] {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if sm.version < len(sm.snapshots)-1 {
		sm.version++
	}

	return sm.snapshots[sm.version]
}
