package commandvalidation

import "fmt"

type ProcessingCommand struct {
	Input string
}

func (p *ProcessingCommand) Execute() (interface{}, error) {
	// Simulate input processing logic here
	result := fmt.Sprintf("Processed: %s", p.Input)
	return result, nil
}
