package tracker

import (
	"fmt"
	"strings"
	"time"

	. "github.com/scotow/skyblocktracker/notifier"
)

func NewEntityTracker(id, name string, interval time.Duration) *EntityTracker {
	et := new(EntityTracker)
	et.id = id
	et.name = name
	et.interval = interval
	return et
}

type EntityTracker struct {
	id         string
	name       string
	interval   time.Duration
	wasPresent bool
}

func (et *EntityTracker) Command() string {
	return fmt.Sprintf("execute if entity @e[type=minecraft:%s] run list", et.id)
}

func (et *EntityTracker) Wait() time.Duration {
	return et.interval
}

func (et *EntityTracker) Track(result string, notifier Notifier) error {
	if len(strings.TrimSpace(result)) > 0 {
		if et.wasPresent {
			return nil
		}

		et.wasPresent = true
		return notifier.Notify(fmt.Sprintf("%s has spawned!", et.name))
	}

	et.wasPresent = false
	return nil
}
