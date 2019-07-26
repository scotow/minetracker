package tracker

import (
	"time"
)

// Tracker is used to build a command string that will be run on the server, parse the result and decide if a notification should be send.
// The Tracker also has to provide an retry/re-run interval.
type Tracker interface {
	// Command, return that the command that will pass to the Runner.
	Command() string

	// Wait calculate and return the amount of time that the Ticker should wait before re-running the command.
	Wait() time.Duration

	// After running the command, the Ticker will pass the result of the command to this Track function.
	// The function should return a bool, indicating if the Ticker should Notify.
	// A string, the message that should be send.
	// An error if something went wrong.
	Track(string) (bool, string, error)
}
