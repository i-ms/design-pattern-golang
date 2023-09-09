package cof_validation

import "errors"

type FormatValidator struct {
	validFormat string
	successor   ValidationHandler
}

func NewFormatValidator(validFormat string) *FormatValidator {
	return &FormatValidator{validFormat: validFormat}
}

func (v *FormatValidator) Validate(input string) error {
	if input != v.validFormat {
		return errors.New("Invalid input format")
	}
	return v.successor.Validate(input)
}

func (v *FormatValidator) SetSuccessor(handler ValidationHandler) {
	v.successor = handler
}
