package abstract_factory_sports

type Nike struct {
}

type NikeShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

func (a *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 9,
		},
	}
}

func (a *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 40,
		},
	}
}
