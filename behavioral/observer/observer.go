package observer

type Publisher interface {
	Update(string, string)
}
