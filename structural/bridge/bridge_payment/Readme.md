# Bridge Pattern

**Architecture Overview**
1. **Abstraction** (`Payment`): Adds features like `Refund()` and `GetStatus()`.
2. **Implementor** (`PaymentMethod`): Adds methods like `Cancel()` and `GetDetails()`.
3. **RefinedAbstraction** (`OnlinePayment`, `OfflinePayment`): Introduce additional types of online and offline payments.
4. **ConcreteImplementor** (`CreditCard`, `PayPal`, `BankTransfer`): Adds more payment methods.

```mermaid
classDiagram
    Payment <|-- OnlinePayment: extends
    Payment <|-- OfflinePayment: extends
    PaymentMethod <|-- CreditCard: implements
    PaymentMethod <|-- PayPal: implements
    PaymentMethod <|-- BankTransfer: implements
    OnlinePayment o-- PaymentMethod: uses
    OfflinePayment o-- PaymentMethod: uses

    class Payment {
        <<interface>>
        +Authorize() string
        +Refund() string
        +GetStatus() string
        +SetPaymentMethod(PaymentMethod)
    }

    class PaymentMethod {
        <<interface>>
        +Process() string
        +Cancel() string
        +GetDetails() string
    }

    class OnlinePayment {
        +Authorize() string
        +Refund() string
        +GetStatus() string
        +SetPaymentMethod(PaymentMethod)
    }

    class OfflinePayment {
        +Authorize() string
        +Refund() string
        +GetStatus() string
        +SetPaymentMethod(PaymentMethod)
    }

    class CreditCard {
        +Process() string
        +Cancel() string
        +GetDetails() string
    }

    class PayPal {
        +Process() string
        +Cancel() string
        +GetDetails() string
    }

    class BankTransfer {
        +Process() string
        +Cancel() string
        +GetDetails() string
    }

```

```mermaid
sequenceDiagram
    participant t as TestCases
    participant op as OnlinePayment
    participant cc as CreditCard
    
    t->>op: SetPaymentMethod(CreditCard)
    t->>op: Authorize()
    op->>cc: Process()
    cc->>op: return "Processed by Credit Card."
    op->>t: return "Online Payment: Processed by Credit Card."
    
    t->>op: Refund()
    op->>cc: Cancel()
    cc->>op: return "Transaction Cancelled by Credit Card."
    op->>t: return "Online Payment Refund: Transaction Cancelled by Credit Card."

    t->>op: GetStatus()
    op->>cc: GetDetails()
    cc->>op: return "Transaction Details from Credit Card."
    op->>t: return "Online Payment Status: Transaction Details from Credit Card."

```

**Explanation**
The current architecture makes it very easy to add new payment methods like stripe, paypal, upi etc. Also, this follows the segregation of principles as the `Payment` interface is not bloated with the methods of `PaymentMethod` interface.