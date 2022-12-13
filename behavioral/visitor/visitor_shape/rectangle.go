package visitor_shape

type Rectangle struct {
	length  int
	breadth int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}
