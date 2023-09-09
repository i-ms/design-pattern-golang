package cof_validation

import "errors"

type LengthValidator struct {
	minLength int
	successor ValidationHandler
}

func NewLengthValidator(minLength int) *LengthValidator {
	return &LengthValidator{minLength: minLength}
}

func (v *LengthValidator) Validate(input string) error {
	if len(input) < v.minLength {
		return errors.New("Input length is too short")
	}
	return v.successor.Validate(input)
}

func (v *LengthValidator) SetSuccessor(handler ValidationHandler) {
	v.successor = handler
}
