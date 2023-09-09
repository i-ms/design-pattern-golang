package cof_validation

type LogErrors struct {
	successor ValidationHandler
}

func NewLogErrors() *LogErrors {
	return &LogErrors{}
}

func (l *LogErrors) Validate(input string) error {
	// Simulate logging validation errors
	println("Validaiton error occured: ", input)
	return nil
}

func (l *LogErrors) SetSuccessor(handler ValidationHandler) {
	l.successor = handler
}
