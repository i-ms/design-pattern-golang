package observer_subject

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
