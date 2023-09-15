package commandvalidation

import "errors"

type InputValidationCommand struct {
	Input string
}

func (i *InputValidationCommand) Execute() (interface{}, error) {
	if len(i.Input) < 5 {
		return false, errors.New("Input must bre greater than 5 characters")
	}
	return true, nil
}
