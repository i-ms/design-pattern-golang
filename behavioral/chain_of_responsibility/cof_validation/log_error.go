package cof_validation

type LogErrors struct {
	SuccessorHandler
}

func NewLogErrors() *LogErrors {
	return &LogErrors{}
}

func (l *LogErrors) Validate(input string) error {
	// Simulate logging validation errors
	println("Validaiton error occured: ", input)
	if l.successor == nil {
		return nil
	}
	return nil
}
