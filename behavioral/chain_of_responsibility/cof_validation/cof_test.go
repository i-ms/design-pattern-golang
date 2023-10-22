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

	// Process the input
	err := processor.ProcessInput("abc123")
	if err != nil {
		println("Input processing failed: ", err)
	} else {
		println("Input processing successed.")
	}
}
