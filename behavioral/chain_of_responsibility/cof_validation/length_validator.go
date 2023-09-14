package cof_validation

import "errors"

type LengthValidator struct {
	minLength int
	SuccessorHandler
}

func NewLengthValidator(minLength int) *LengthValidator {
	return &LengthValidator{
		minLength: minLength,
	}
}

func (v *LengthValidator) Validate(input string) error {
	if len(input) < v.minLength {
		return errors.New("Input length is too short")
	}
	if v.successor == nil {
		return nil
	}
	return v.successor.Validate(input)
}
