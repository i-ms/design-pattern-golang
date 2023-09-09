package cof_validation

import "testing"

func TestCofValidation(t *testing.T) {
	processor := NewInputProcessor()
	processor.SetHandlers(
		NewLengthValidator(5),
		NewFormatValidator("abc123"),
		NewSanitizeInput(),
		NewStoreInDatabase(),
		NewLogErrors(),
	)

	// Simulate user input
	input := "abc123"

	// Process the input
	err := processor.ProcessInput(input)
	if err != nil {
		println("Input processing failed: ", err)
	} else {
		println("Input processing successed.")
	}
}
