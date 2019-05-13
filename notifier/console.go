package notifier

import (
	"fmt"
)

func NewConsoleNotifier() *ConsoleNotifier {
	return new(ConsoleNotifier)
}

type ConsoleNotifier struct{}

func (cn *ConsoleNotifier) Notify(data string) error {
	fmt.Println(data)
	return nil
}
