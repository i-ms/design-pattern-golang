package cof_validation

type SanitizeInput struct {
	SuccessorHandler
}

func NewSanitizeInput() *SanitizeInput {
	return &SanitizeInput{}
}

func (s *SanitizeInput) Validate(input string) error {
	if s.successor == nil {
		return nil
	}
	// Perform input sanitization here (e.g. remove all non-alphanumeric characters)
	// For simplicity, we'll skip the actual sanitization step
	return s.successor.Validate(input)
}
