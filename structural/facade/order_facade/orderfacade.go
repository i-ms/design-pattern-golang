package orderfacade

import "fmt"

type OrderFacade struct {
	inventory        *Inventory
	shoppingCart     *ShoppingCart
	paymentProcessor *PaymentProcessor
	shippingService  *ShippingService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory:        NewInventory(),
		shoppingCart:     NewShoppingCart(),
		paymentProcessor: &PaymentProcessor{},
		shippingService:  &ShippingService{},
	}
}

func (of *OrderFacade) AddProductToCart(productID, quantity int) {
	product, exists := of.inventory.GetProduct(productID)
	if !exists {
		println("Product not found")
		return
	}

	of.shoppingCart.AddToCart(productID, quantity)
	fmt.Printf("Added %d %s to the shopping cart.\n", quantity, product.Name)
}

func (of *OrderFacade) Checkout() {
	cartContents := of.shoppingCart.GetCartContents()
	var totalAmount float64

	for productID, quantity := range cartContents {
		product, exists := of.inventory.GetProduct(productID)
		if !exists {
			println("Product not found")
			continue
		}
		totalAmount += product.Price * float64(quantity)
	}

	if of.paymentProcessor.ProcessPayment(totalAmount) {
		fmt.Printf("Payment successfull. Total amount: %f\n", totalAmount)
		of.shippingService.ShipProducts(cartContents)
		of.shoppingCart = NewShoppingCart()
	} else {
		fmt.Println("Payment failed. Please try again.")
	}
}
