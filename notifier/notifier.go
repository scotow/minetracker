package notifier

// A Notifier is used to send the information about a success tracking event.
type Notifier interface {
	// Send the following string.
	Notify(string) error
}
