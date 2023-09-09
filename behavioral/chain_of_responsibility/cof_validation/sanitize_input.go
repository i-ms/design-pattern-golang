package cof_validation

type SanitizeInput struct {
	successor ValidationHandler
}

func NewSanitizeInput() *SanitizeInput {
	return &SanitizeInput{}
}

func (s *SanitizeInput) Validate(input string) error {
	// Perform input sanitization here (e.g. remove all non-alphanumeric characters)
	// For simplicity, we'll skip the actual sanitization step
	return s.successor.Validate(input)
}

func (s *SanitizeInput) SetSuccessor(handler ValidationHandler) {
	s.successor = handler
}
