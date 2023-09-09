package cof_validation

type StoreInDatabase struct {
	successor ValidationHandler
}

func NewStoreInDatabase() *StoreInDatabase {
	return &StoreInDatabase{}
}

func (s *StoreInDatabase) Validate(input string) error {
	// Simulate storing input in database
	println("Input data stored in database")
	return s.successor.Validate(input)
}

func (s *StoreInDatabase) SetSuccessor(handler ValidationHandler) {
	s.successor = handler
}
