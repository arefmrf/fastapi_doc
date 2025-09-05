package main

import "fmt"

// Logger is the expected interface in our application.
type Logger interface {
	Log(message string)
}

// ThirdPartyLogger is a library that logs messages differently.
type ThirdPartyLogger struct{}

// PrintMessage is the third-party method (does not match Logger interface).
func (t *ThirdPartyLogger) PrintMessage(msg string) {
	fmt.Println("Third-Party Logger:", msg)
}

type LoggerAdapter struct {
	thirdPartyLogger *ThirdPartyLogger
}

// Log adapts the PrintMessage method to match the Logger interface.
func (l *LoggerAdapter) Log(message string) {
	l.thirdPartyLogger.PrintMessage(message) // ✅ Calls the third-party method
}

func Adapt() {
	// Create an instance of the third-party logger
	thirdParty := &ThirdPartyLogger{}

	// Use the adapter to make it compatible with our Logger interface
	adapter := &LoggerAdapter{thirdPartyLogger: thirdParty}

	// Call the Log method (which internally calls PrintMessage)
	adapter.Log("Hello, Adapter Pattern!")
}

//
// ****************** second example ********************************
//

type OldPrinter struct{}

func (op *OldPrinter) PrintOld(msg string) {
	fmt.Println("Old Printer: " + msg)
}

// ModernPrinter is the interface we want to use
type ModernPrinter interface {
	Print(msg string)
}

type PrinterAdapter struct {
	OldPrinter *OldPrinter
}

// Implement ModernPrinter's Print method
func (pa *PrinterAdapter) Print(msg string) {
	pa.OldPrinter.PrintOld(msg) // Calls the old method
}

func AdaptPrint() {
	oldPrinter := &OldPrinter{}                        // Existing old printer
	adapter := &PrinterAdapter{OldPrinter: oldPrinter} // Create adapter

	// Use adapter as a ModernPrinter
	var printer ModernPrinter = adapter
	printer.Print("Hello, world!") // Calls PrintOld() internally
}
