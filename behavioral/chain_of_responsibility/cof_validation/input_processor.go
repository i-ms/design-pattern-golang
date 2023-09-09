package cof_validation

type InputProcessor struct {
	handlers []ValidationHandler
}

func NewInputProcessor() *InputProcessor {
	return &InputProcessor{}
}

func (p *InputProcessor) SetHandlers(handlers ...ValidationHandler) {
	for i := 0; i < len(handlers)-1; i++ {
		handlers[i].SetSuccessor(handlers[i+1])
	}
	p.handlers = handlers
}

func (p *InputProcessor) ProcessInput(input string) error {
	for _, handler := range p.handlers {
		if err := handler.Validate(input); err != nil {
			return err
		}
	}
	return nil
}
