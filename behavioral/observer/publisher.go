package observer

type Publisher interface {
	CreatePost(string)
	Notify(string)
}
