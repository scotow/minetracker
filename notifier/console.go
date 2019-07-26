package notifier

import (
	"fmt"
)

// Create a Console Notifier.
func NewConsoleNotifier() *ConsoleNotifier {
	return new(ConsoleNotifier)
}

// A simple Notifier that prints the message to stdout.
type ConsoleNotifier struct{}

func (cn *ConsoleNotifier) Notify(data string) error {
	fmt.Println(data)
	return nil
}
