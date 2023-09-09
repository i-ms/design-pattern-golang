package orderfacade

type ShoppingCart struct {
	contents map[int]int
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		contents: make(map[int]int),
	}
}

func (s *ShoppingCart) AddToCart(id int, quantity int) {
	s.contents[id] = quantity
}

func (s *ShoppingCart) GetCartContents() map[int]int {
	return s.contents
}
