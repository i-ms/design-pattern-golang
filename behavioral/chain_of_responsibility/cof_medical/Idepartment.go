package cof_medical

type Department interface {
	execute(*Patient)
	setNext(Department)
}
