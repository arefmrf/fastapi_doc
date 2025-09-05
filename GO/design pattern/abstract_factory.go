package main

import "fmt"

// Notification is the common interface
type Notification interface {
	SendNotification(message string)
}

// EmailNotification struct
type EmailNotification struct{}

func (e *EmailNotification) SendNotification(message string) {
	fmt.Println("Sending Email:", message)
}

// SMSNotification struct
type SMSNotification struct{}

func (s *SMSNotification) SendNotification(message string) {
	fmt.Println("Sending SMS:", message)
}

// NotificationFactory is an interface for factories
type NotificationFactory interface {
	CreateNotification() Notification
}

// EmailFactory struct
type EmailFactory struct{}

func (e *EmailFactory) CreateNotification() Notification {
	return &EmailNotification{}
}

// SMSFactory struct
type SMSFactory struct{}

func (s *SMSFactory) CreateNotification() Notification {
	return &SMSNotification{}
}

func RunAbstractFactory() {
	// Using the factory to create an email notification
	emailFactory := &EmailFactory{}
	email := emailFactory.CreateNotification()
	email.SendNotification("Hello via Email!")

	// Using the factory to create an SMS notification
	smsFactory := &SMSFactory{}
	sms := smsFactory.CreateNotification()
	sms.SendNotification("Hello via SMS!")
}
