package iterator_collection

type Collection interface {
	createIterator() Iterator
}
