package abstract_factory_sports

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "Adidas",
			size: 9,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "Adidas",
			size: 40,
		},
	}
}
