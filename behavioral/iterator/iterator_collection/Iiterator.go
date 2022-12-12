package iterator_collection

type Iterator interface {
	hasNext() bool
	getNext() *User
}
