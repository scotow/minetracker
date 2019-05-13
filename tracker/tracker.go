package tracker

import (
	. "github.com/scotow/skyblocktracker/notifier"
)

type Tracker interface {
	Command() string
	Track(string, Notifier) error
}
