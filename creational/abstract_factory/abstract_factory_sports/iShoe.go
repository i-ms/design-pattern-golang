package abstract_factory_sports

type IShoe interface {
	setLogo(string)
	setSize(int)
	getLogo() string
	getSize() int
}
