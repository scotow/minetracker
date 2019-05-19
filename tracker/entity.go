package tracker

import (
	"fmt"
	"strings"
	"time"
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
	return fmt.Sprintf("execute if entity @e[type=minecraft:%s] run time query daytime", et.id)
}

func (et *EntityTracker) Wait() time.Duration {
	return et.interval
}

func (et *EntityTracker) Track(result string) (bool, string, error) {
	if len(strings.TrimSpace(result)) > 0 {
		if et.wasPresent {
			return false, "", nil
		}

		et.wasPresent = true
		return true, fmt.Sprintf("%s has spawned!", et.name), nil
	}

	et.wasPresent = false
	return false, "", nil
}
