package orderfacade

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Inventory struct {
	products map[int]Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[int]Product),
	}
}

func (i *Inventory) AddProduct(id int, name string, price float64) {
	i.products[id] = Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

func (i *Inventory) GetProduct(id int) (Product, bool) {
	product, exists := i.products[id]
	return product, exists
}
