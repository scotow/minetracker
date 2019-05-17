package tracker

import (
	"time"
)

type Tracker interface {
	Command() string
	Wait() time.Duration
	Track(string) (bool, string, error)
}
