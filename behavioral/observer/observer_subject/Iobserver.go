package observer_subject

type Observer interface {
	update(string)
	getID() string
}
