package tracker

import (
	"time"

	. "github.com/scotow/skyblocktracker/notifier"
)

type Tracker interface {
	Command() string
	Wait() time.Duration
	Track(string, Notifier) error
}
