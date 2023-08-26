# Abstract Factory Pattern

## Introduction
This documentation provides an overview of an Abstract Factory pattern implemented for a Payment System in Go. The system is designed to handle two types of payment accounts: PayPal and Stripe, each having two subtypes: Personal and Business accounts.

## Architecture
The system is structured around two main interfaces:

- **AccountFactory**: An interface for creating different types of payment accounts.
- **PaymentMethod**: An interface defining various account operations like Pay, Refund, and Transfer.

There are two concrete implementations for AccountFactory:
- **PaypalAccountFactory**: Creates PayPal Personal and Business accounts.
- **StripeAccountFactory**: Creates Stripe Personal and Business accounts.
Each account type (PayPal or Stripe, Personal or Business) implements the PaymentMethod interface.

```mermaid
classDiagram
    class AccountFactory {
        <<interface>>
        +CreatePersonalAccount() PaymentMethod
        +CreateBusinessAccount() PaymentMethod
    }
    class PaymentMethod {
        <<interface>>
        +Pay(float64) string
        +Refund(float64) string
        +Transfer(PaymentMethod, float64) string
    }
    class PaypalAccountFactory {
        +CreatePersonalAccount() PaymentMethod
        +CreateBusinessAccount() PaymentMethod
    }
    class StripeAccountFactory {
        +CreatePersonalAccount() PaymentMethod
        +CreateBusinessAccount() PaymentMethod
    }
    class PaypalPersonalAccount {
        +Pay(float64) string
        +Refund(float64) string
        +Transfer(PaymentMethod, float64) string
    }
    class PaypalBusinessAccount {
        +Pay(float64) string
        +Refund(float64) string
        +Transfer(PaymentMethod, float64) string
    }
    class StripePersonalAccount {
        +Pay(float64) string
        +Refund(float64) string
        +Transfer(PaymentMethod, float64) string
    }
    class StripeBusinessAccount {
        +Pay(float64) string
        +Refund(float64) string
        +Transfer(PaymentMethod, float64) string
    }
    AccountFactory --|> PaypalAccountFactory
    AccountFactory --|> StripeAccountFactory
    PaymentMethod --|> PaypalPersonalAccount
    PaymentMethod --|> PaypalBusinessAccount
    PaymentMethod --|> StripePersonalAccount
    PaymentMethod --|> StripeBusinessAccount

```

```mermaid
sequenceDiagram
    participant Client
    participant AccountFactory
    participant PaymentMethod
    participant PaypalAccountFactory
    participant StripeAccountFactory
    participant PaypalPersonalAccount
    participant PaypalBusinessAccount
    participant StripePersonalAccount
    participant StripeBusinessAccount

    Client->>AccountFactory: Request CreatePersonalAccount()
    AccountFactory-->>PaypalAccountFactory: CreatePersonalAccount()
    AccountFactory-->>StripeAccountFactory: CreatePersonalAccount()
    PaypalAccountFactory-->>PaymentMethod: Implement PaymentMethod
    StripeAccountFactory-->>PaymentMethod: Implement PaymentMethod
    PaymentMethod-->>PaypalPersonalAccount: Pay(100)
    PaymentMethod-->>PaypalBusinessAccount: Pay(200)
    PaymentMethod-->>StripePersonalAccount: Pay(300)
    PaymentMethod-->>StripeBusinessAccount: Pay(400)
```
## Explanation
The AccountFactory interface provides methods to create Personal and Business accounts. The concrete implementations like PaypalAccountFactory and StripeAccountFactory implement these methods and return instances of accounts that implement the PaymentMethod interface. The PaymentMethod interface has methods for conducting payments, refunds, and transfers. Each concrete account type (Personal and Business for both PayPal and Stripe) implements these methods.