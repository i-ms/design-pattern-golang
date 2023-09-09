package cof_validation

type ValidationHandler interface {
	Validate(input string) error
	SetSuccessor(handler ValidationHandler)
}
