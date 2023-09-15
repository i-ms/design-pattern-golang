package commandvalidation

type Invoker struct{}

func (i *Invoker) Execute(cmd Command) (interface{}, error) {
	return cmd.Execute()
}
