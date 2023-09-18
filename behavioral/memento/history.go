package memento

type History struct {
	snapshotManager *SnapshotManager[CodeEditor]
}

func newHisory() *History {
	return &History{
		snapshotManager: NewSnapshotManager[CodeEditor](),
	}
}

func (h *History) save(c CodeEditor) {
	h.snapshotManager.AddSnapshot(Snapshot[CodeEditor]{
		State: c,
	})
}

func (h *History) undo() Snapshot[CodeEditor] {
	return h.snapshotManager.GetPreviousSnapshot()
}

func (h *History) redo() Snapshot[CodeEditor] {
	return h.snapshotManager.GetNextSnapshot()
}
