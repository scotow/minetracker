package tracker

import (
	"fmt"
	"strings"
	"time"

	. "github.com/scotow/skyblocktracker/notifier"
)

func NewEntityTracker(id, name string, interval, wait time.Duration) *EntityTracker {
	et := new(EntityTracker)
	et.id = id
	et.name = name
	et.interval = interval
	et.wait = wait
	return et
}

type EntityTracker struct {
	id             string
	name           string
	interval, wait time.Duration
	isLongInterval bool
}

func (et *EntityTracker) Command() string {
	return fmt.Sprintf("execute if entity @e[type=minecraft:%s] run list", et.id)
}

func (et *EntityTracker) Wait() time.Duration {
	if et.isLongInterval {
		et.isLongInterval = false
		return et.wait
	} else {
		return et.interval
	}
}

func (et *EntityTracker) Track(result string, notifier Notifier) error {
	if len(strings.TrimSpace(result)) > 0 {
		et.isLongInterval = true
		return notifier.Notify(fmt.Sprintf("%s has spawned!", et.name))
	}

	return nil
}
