package memento

import "testing"

func TestCodeEditor(t *testing.T) {
	codeEditor := NewCodeEditor()
	codeEditor.Write("Hello")
	codeEditor.Write(" World")
	codeEditor.undo()
	codeEditor.redo()
}
