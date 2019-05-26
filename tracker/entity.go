package tracker

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrInvalidTestCommand = errors.New("invalid command result for entity testing")
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
	trimmed := strings.TrimSpace(result)

	if len(trimmed) == 0 {
		et.wasPresent = false
		return false, "", nil
	}

	n, err := fmt.Sscanf(trimmed, "The time is %d.", new(int))
	if n != 1 || err != nil {
		return false, "", ErrInvalidTestCommand
	}

	if et.wasPresent {
		return false, "", nil
	}

	et.wasPresent = true
	return true, fmt.Sprintf("%s has spawned!", et.name), nil
}
