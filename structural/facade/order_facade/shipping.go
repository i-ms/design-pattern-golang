package orderfacade

type ShippingService struct{}

func (ss *ShippingService) ShipProducts(products map[int]int) {
	println("Shipping products : ", products)
}
