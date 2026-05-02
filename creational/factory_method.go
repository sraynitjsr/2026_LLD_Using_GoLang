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

// sendNotification demonstrates using the factory method
func sendNotification(factory NotificationFactory, message string) {
	notification := factory.CreateNotification()
	notification.Send(message)
}

// MyFactoryMethod demonstrates the factory method pattern
func MyFactoryMethod() {
	fmt.Println("Factory Method Pattern: Creating different notification types")
	fmt.Println()

	// Create factories
	emailFactory := &EmailNotificationFactory{recipient: "user@example.com"}
	smsFactory := &SMSNotificationFactory{phoneNumber: "+1-555-1234"}
	pushFactory := &PushNotificationFactory{deviceID: "device-abc123"}

	// Use the same function to send different types of notifications
	sendNotification(emailFactory, "Welcome to our service!")
	sendNotification(smsFactory, "Your verification code is 123456")
	sendNotification(pushFactory, "You have a new message")

	fmt.Println("\nFactory Method: Define interface for creating objects, let subclasses decide which class to instantiate")
}
