# Facade Pattern

### Class Diagram

```mermaid
classDiagram
  class Product {
    +ID: int
    +Name: string
    +Price: float64
    +AddProduct(id: int, name: string, price: float64)
    +GetProduct(id: int): (Product, bool)
  }

  class Inventory {
    -products: map[int]Product
    +AddProduct(id: int, name: string, price: float64)
    +GetProduct(id: int): (Product, bool)
  }

  class ShoppingCart {
    -contents: map[int]int
    +AddToCart(productID: int, quantity: int)
    +GetCartContents(): map[int]int
  }

  class PaymentProcessor {
    +ProcessPayment(amount: float64): bool
  }

  class ShippingService {
    +ShipProducts(products: map[int]int)
  }

  class OrderFacade {
    -inventory: Inventory
    -shoppingCart: ShoppingCart
    -paymentProcessor: PaymentProcessor
    -shippingService: ShippingService
    +NewOrderFacade(): OrderFacade
    +AddProductToCart(productID: int, quantity: int)
    +Checkout()
  }

  Product --* Inventory: Contains
  ShoppingCart --* Product: Contains
  OrderFacade o-- Inventory: Uses
  OrderFacade o-- ShoppingCart: Uses
  OrderFacade o-- PaymentProcessor: Uses
  OrderFacade o-- ShippingService: Uses
```

### Sequence Diagram

```mermaid
sequenceDiagram
  participant Client
  participant OrderFacade
  participant Inventory
  participant ShoppingCart
  participant PaymentProcessor
  participant ShippingService

  Client->>OrderFacade: NewOrderFacade()
  Note over OrderFacade: Creates instances of subsystems\nand initializes them.

  Client->>OrderFacade: AddProductToCart(1, 2)
  OrderFacade->>Inventory: GetProduct(1)
  Inventory-->>OrderFacade: Product(1, "Laptop", 800.00)
  OrderFacade->>ShoppingCart: AddToCart(1, 2)
  ShoppingCart->>ShoppingCart: Add 2 Laptops to cart

  Client->>OrderFacade: AddProductToCart(2, 3)
  OrderFacade->>Inventory: GetProduct(2)
  Inventory-->>OrderFacade: Product(2, "Headphones", 50.00)
  OrderFacade->>ShoppingCart: AddToCart(2, 3)
  ShoppingCart->>ShoppingCart: Add 3 Headphones to cart

  Client->>OrderFacade: Checkout()
  OrderFacade->>ShoppingCart: GetCartContents()
  ShoppingCart-->>OrderFacade: Cart Contents
  OrderFacade->>Inventory: GetProduct(1)
  Inventory-->>OrderFacade: Product(1, "Laptop", 800.00)
  OrderFacade->>Inventory: GetProduct(2)
  Inventory-->>OrderFacade: Product(2, "Headphones", 50.00)
  OrderFacade->>PaymentProcessor: ProcessPayment(2750.00)
  PaymentProcessor-->>OrderFacade: true
  OrderFacade->>ShippingService: ShipProducts(Cart Contents)
  ShippingService-->>OrderFacade: Shipping products: {1: 2, 2: 3}
  OrderFacade->>ShoppingCart: NewShoppingCart()
  ShoppingCart-->>OrderFacade: Cart Emptied

```
