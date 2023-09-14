package cof_validation

import "errors"

type FormatValidator struct {
	validFormat string
	SuccessorHandler
}

func NewFormatValidator(validFormat string) *FormatValidator {
	return &FormatValidator{validFormat: validFormat}
}

func (v *FormatValidator) Validate(input string) error {
	if input != v.validFormat {
		return errors.New("Invalid input format")
	}
	if v.successor == nil {
		return nil
	}
	return v.successor.Validate(input)
}
