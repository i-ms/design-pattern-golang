package commandvalidation

type Command interface {
	Execute() (interface{}, error)
}
