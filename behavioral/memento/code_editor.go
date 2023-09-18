package memento

type CodeEditor struct {
	content string
	history History
}

func NewCodeEditor() *CodeEditor {
	return &CodeEditor{
		history: *newHisory(),
	}
}

func (c *CodeEditor) Write(content string) {
	c.content += content
	println("Current Editor State:", c.content)
	c.history.save(*c)
}

func (c *CodeEditor) undo() {
	println("Editor State before undo:", c.content)
	snapshot := c.history.undo()
	c.content = snapshot.State.content
	println("Editor State after undo:", c.content)
}

func (c *CodeEditor) redo() {
	println("Editor State before redo:", c.content)
	snapshot := c.history.redo()
	c.content = snapshot.State.content
	println("Editor State after redo:", c.content)
}
