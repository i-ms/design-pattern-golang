# Adapter Pattern

## Architecture 
The architecture primarily consists of the following components:
1. `Payment` : THe target interface to standardise payment authorization.
2. `CredtCard` and `Paypal` : The adaptees which contains different APIs for processing payment.
3. `CredtCardAdapter` and `PaypalAdapter` : The adapters that make `CreditCard` and `Paypal` conform to the Payment interface.

## Concept
The Adapter Pattern allows classes with incompatible interfaces to work together. In this example, we  have two different payment systems(`CreditCard` and `Paypal`) with their APIs (`Validate()` and `Confirm()`).The `Adapter` struct (`CreditCardAdapter` and `PaypalAdapter`) make these different systems compatible with a common interface `Payment`. 

```mermaid
classDiagram
    Payment <|-- PayPalAdapter : implements
    Payment <|-- CreditCardAdapter : implements
    PayPalAdapter o-- PayPal : adapts
    CreditCardAdapter o-- CreditCard : adapts

    class Payment {
        <<interface>>
        +Authorize() string
    }

    class PayPal {
        +Confirm() string
    }

    class PayPalAdapter {
        +Authorize() string
    }

    class CreditCard {
        +Validate() string
    }

    class CreditCardAdapter {
        +Authorize() string
    }
```

```mermaid
sequenceDiagram
    participant t as TestCases
    participant pp as ProcessPayment()
    participant cca as CreditCardAdapter
    participant cc as CreditCard
    participant ppa as PayPalAdapter
    participant pp as PayPal

    t->>pp: call ProcessPayment()
    pp->>cca: call Authorize()
    cca->>cc: call Validate()
    cc->>cca: return "Credit Card payment validated."
    cca->>pp: return "Credit Card payment validated."
    pp->>t: return
    
    t->>pp: call ProcessPayment()
    pp->>ppa: call Authorize()
    ppa->>pp: call Confirm()
    pp->>ppa: return "PayPal payment confirmed."
    ppa->>pp: return "PayPal payment confirmed."
    pp->>t: return
```