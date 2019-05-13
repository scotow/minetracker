package tracker

import (
	"fmt"
	"strings"

	. "github.com/scotow/skyblocktracker/notifier"
)

func NewEntityTracker(id, name string) *EntityTracker {
	et := new(EntityTracker)
	et.id = id
	et.name = name
	return et
}

type EntityTracker struct {
	id   string
	name string
}

func (et *EntityTracker) Command() string {
	return fmt.Sprintf("execute if entity @e[type=minecraft:%s] run list", et.id)
}

func (et *EntityTracker) Track(result string, notifier Notifier) error {
	if len(strings.TrimSpace(result)) > 0 {
		return notifier.Notify(fmt.Sprintf("%s has spawned!", et.name))
	}

	return nil
}
