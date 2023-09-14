package cof_validation

type StoreInDatabase struct {
	SuccessorHandler
}

func NewStoreInDatabase() *StoreInDatabase {
	return &StoreInDatabase{}
}

func (s *StoreInDatabase) Validate(input string) error {
	// Simulate storing input in database
	println("Input data stored in database")
	if s.successor == nil {
		return nil
	}
	return s.successor.Validate(input)
}
