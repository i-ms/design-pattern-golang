package cof_validation

type SuccessorHandler struct {
	successor ValidationHandler
}

func (s *SuccessorHandler) SetSuccessor(handler ValidationHandler) {
	s.successor = handler
}
