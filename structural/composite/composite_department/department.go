package composite_department

import (
	"fmt"
)

type Component interface {
	Display(indent int)
}

type Employee struct {
	Name     string
	Position string
	Salary   float64
}

func (e *Employee) Display(indent int) {
	_ = fmt.Sprintf(
		"%sEmployee Information\nName: %s\nPosition: %s\nSalary: %f",
		getIndent(indent),
		e.Name,
		e.Position,
		e.Salary,
	)
}

// Department represents  a department within the organisation
type Department struct {
	Name      string
	Component []Component
}

func (d *Department) Add(component Component) {
	d.Component = append(d.Component, component)
}

func (d *Department) Display(indent int) {
	_ = fmt.Sprintf("%sDepartment: %s\n", getIndent(indent), d.Name)
	for _, component := range d.Component {
		component.Display(indent + 2)
	}
}

func getIndent(indent int) string {
	return string(make([]byte, indent))
}
