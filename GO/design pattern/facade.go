package main

import "fmt"

// Light component
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is turned ON")
}

func (l *Light) Off() {
	fmt.Println("Light is turned OFF")
}

// TV component
type TV struct{}

func (tv *TV) On() {
	fmt.Println("TV is turned ON")
}

func (tv *TV) Off() {
	fmt.Println("TV is turned OFF")
}

// AirConditioner component
type AirConditioner struct{}

func (ac *AirConditioner) On() {
	fmt.Println("Air Conditioner is turned ON")
}

func (ac *AirConditioner) Off() {
	fmt.Println("Air Conditioner is turned OFF")
}

// ----
// ----
// HomeAutomationFacade simplifies operations
type HomeAutomationFacade struct {
	light          *Light
	tv             *TV
	airConditioner *AirConditioner
}

// NewHomeAutomationFacade creates a new facade instance
func NewHomeAutomationFacade() *HomeAutomationFacade {
	return &HomeAutomationFacade{
		light:          &Light{},
		tv:             &TV{},
		airConditioner: &AirConditioner{},
	}
}

// TurnEverythingOn turns all devices ON
func (h *HomeAutomationFacade) TurnEverythingOn() {
	fmt.Println("Turning everything ON...")
	h.light.On()
	h.tv.On()
	h.airConditioner.On()
}

// TurnEverythingOff turns all devices OFF
func (h *HomeAutomationFacade) TurnEverythingOff() {
	fmt.Println("Turning everything OFF...")
	h.light.Off()
	h.tv.Off()
	h.airConditioner.Off()
}

// ----
// ----
func Facade() {
	home := NewHomeAutomationFacade()

	// Turn everything ON
	home.TurnEverythingOn()

	fmt.Println()

	// Turn everything OFF
	home.TurnEverythingOff()
}
