package orderfacade

import "testing"

func TestOrderFacade(t *testing.T) {
	orderFacade := NewOrderFacade()

	// Adding products to inventory
	orderFacade.inventory.AddProduct(1, "Motherboard", 100)
	orderFacade.inventory.AddProduct(2, "CPU", 200)
	orderFacade.inventory.AddProduct(3, "Ram", 200)
	// Adding products to the shopping cart
	orderFacade.AddProductToCart(1, 2)
	orderFacade.AddProductToCart(2, 1)

	// Checkout the shopping cart
	orderFacade.Checkout()
}
