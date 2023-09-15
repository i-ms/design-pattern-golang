package commandvalidation

import "testing"

func TestCommandValidation(t *testing.T) {
	invoker := Invoker{}

	// Define user input
	userInput := "Hello World"

	// Create validation and processing command
	validation := &InputValidationCommand{Input: userInput}
	processing := &ProcessingCommand{Input: userInput}

	// Execute the validation command
	isValid, validationErr := invoker.Execute(validation)
	if !isValid.(bool) {
		println("Validation failed: ", validationErr.Error())
	}

	// Execute the processing command
	processingResult, processingErr := invoker.Execute(processing)
	if processingErr != nil {
		println("Processing  error: ", processingErr.Error())
		return
	}
	println(processingResult.(string))
}
