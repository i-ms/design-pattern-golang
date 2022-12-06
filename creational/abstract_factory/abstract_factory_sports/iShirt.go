package abstract_factory_sports

type IShirt interface {
	setLogo(string)
	setSize(int)
	getLogo() string
	getSize() int
}
