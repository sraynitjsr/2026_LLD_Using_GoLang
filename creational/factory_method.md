# Factory Method Design Pattern

## Overview
The Factory Method is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. It defines an interface for creating an object but lets subclasses decide which class to instantiate.

## Intent
- Define an interface for creating an object
- Let subclasses decide which class to instantiate
- Defer instantiation to subclasses
- Promote loose coupling by eliminating the need to bind application-specific classes into code

## When to Use
- **Unknown Object Types**: When you don't know beforehand the exact types and dependencies of the objects your code should work with
- **Extension Points**: When you want to provide users of your library or framework with a way to extend its internal components
- **Reuse Existing Objects**: When you want to save system resources by reusing existing objects instead of rebuilding them each time
- **Plugin Architecture**: When you need to support different implementations of the same interface
- **Testing**: When you want to easily substitute implementations for testing purposes

## Real-World Use Cases
- **Notification Systems**: Creating different types of notifications (Email, SMS, Push)
- **Document Generators**: Creating different document formats (PDF, Word, HTML)
- **Database Connections**: Creating connections to different database types (MySQL, PostgreSQL, MongoDB)
- **Payment Processors**: Creating different payment method handlers (Credit Card, PayPal, Crypto)
- **Logistics**: Creating different types of transport (Truck, Ship, Plane)
- **UI Components**: Creating platform-specific UI elements (Windows, macOS, Linux)

## Implementation in Go

This implementation demonstrates a notification system using the factory method pattern:

```go
package creational

import "fmt"

// Notification is the product interface
type Notification interface {
	Send(message string)
}

// EmailNotification is a concrete product
type EmailNotification struct {
	recipient string
}

func (e *EmailNotification) Send(message string) {
	fmt.Printf("📧 Email sent to %s: %s\n", e.recipient, message)
}

// SMSNotification is a concrete product
type SMSNotification struct {
	phoneNumber string
}

func (s *SMSNotification) Send(message string) {
	fmt.Printf("📱 SMS sent to %s: %s\n", s.phoneNumber, message)
}

// PushNotification is a concrete product
type PushNotification struct {
	deviceID string
}

func (p *PushNotification) Send(message string) {
	fmt.Printf("🔔 Push notification sent to device %s: %s\n", p.deviceID, message)
}

// NotificationFactory is the creator interface
type NotificationFactory interface {
	CreateNotification() Notification
}

// EmailNotificationFactory is a concrete creator
type EmailNotificationFactory struct {
	recipient string
}

func (f *EmailNotificationFactory) CreateNotification() Notification {
	return &EmailNotification{recipient: f.recipient}
}

// SMSNotificationFactory is a concrete creator
type SMSNotificationFactory struct {
	phoneNumber string
}

func (f *SMSNotificationFactory) CreateNotification() Notification {
	return &SMSNotification{phoneNumber: f.phoneNumber}
}

// PushNotificationFactory is a concrete creator
type PushNotificationFactory struct {
	deviceID string
}

func (f *PushNotificationFactory) CreateNotification() Notification {
	return &PushNotification{deviceID: f.deviceID}
}
```

## Key Components

### 1. Product Interface (Notification)
Declares the interface that all concrete products must implement.

### 2. Concrete Products (EmailNotification, SMSNotification, PushNotification)
Different implementations of the product interface.

### 3. Creator Interface (NotificationFactory)
Declares the factory method that returns new product objects.

### 4. Concrete Creators (EmailNotificationFactory, SMSNotificationFactory, PushNotificationFactory)
Override the factory method to return different types of products.

## Benefits
- **Open/Closed Principle**: You can introduce new types of products without breaking existing client code
- **Single Responsibility Principle**: You can move the product creation code into one place
- **Loose Coupling**: The code only works with the interface, not concrete classes
- **Flexibility**: Easy to add new product types by creating new concrete creators
- **Testability**: Easy to mock and test different implementations

## Drawbacks
- **Complexity**: The code may become more complicated since you need to introduce multiple new classes
- **Hierarchy**: Requires creating a parallel class hierarchy of creators and products

## Comparison with Other Patterns

### Factory Method vs Abstract Factory
- **Factory Method**: Single method to create one type of product
- **Abstract Factory**: Multiple methods to create families of related products

### Factory Method vs Simple Factory
- **Factory Method**: Uses inheritance and delegates object creation to subclasses
- **Simple Factory**: Uses a single factory class with conditional logic

### Factory Method vs Builder
- **Factory Method**: Creates complete objects in one step
- **Builder**: Constructs complex objects step by step

## Go-Specific Considerations
- Go doesn't have classes, so we use interfaces and structs
- Factory functions can be regular functions or methods on factory structs
- Go's implicit interface satisfaction makes it easy to add new product types
- Consider using functional options pattern for complex product configuration

## Example Usage
```go
// Create factories
emailFactory := &EmailNotificationFactory{recipient: "user@example.com"}
smsFactory := &SMSNotificationFactory{phoneNumber: "+1-555-1234"}

// Use factories to create and send notifications
notification1 := emailFactory.CreateNotification()
notification1.Send("Welcome!")

notification2 := smsFactory.CreateNotification()
notification2.Send("Verification code: 123456")
```

## Related Patterns
- **Abstract Factory**: Often implemented using factory methods
- **Template Method**: Factory methods are often called within template methods
- **Prototype**: Can be used as an alternative when subclassing is not desired
- **Singleton**: Factory methods can ensure singleton instances
