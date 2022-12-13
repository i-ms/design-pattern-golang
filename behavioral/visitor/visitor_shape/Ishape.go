package visitor_shape

type Shape interface {
	getType() string
	accept(Visitor)
}
