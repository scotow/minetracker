package tracker

import (
	"fmt"
	"strings"
	"time"

	. "github.com/scotow/skyblocktracker/misc"
)

func NewConnectionTracker(exclude string, interval time.Duration) *ConnectionTracker {
	ct := new(ConnectionTracker)
	ct.exclude = exclude
	ct.interval = interval
	return ct
}

type ConnectionTracker struct {
	last     []string
	exclude  string
	interval time.Duration
}

func (ct *ConnectionTracker) Command() string {
	return "list"
}

func (ct *ConnectionTracker) Wait() time.Duration {
	return ct.interval
}

func (ct *ConnectionTracker) Track(result string) (bool, string, error) {
	online, err := ParseOnlinePlayers(result)
	if err != nil {
		return false, "", err
	}

	// First track, don't notify.
	if ct.last == nil {
		ct.last = online
		return false, "", nil
	}

	newConnect := FindNew(ct.last, online)
	newDisconnect := FindNew(online, ct.last)

	ct.last = online

	if len(newConnect)+len(newDisconnect) > 0 {
		return ct.excludeAndFormat(online, newConnect, newDisconnect)
	}

	return false, "", nil
}

func (ct *ConnectionTracker) excludeAndFormat(online, connect, disconnect []string) (bool, string, error) {
	if ct.exclude != "" {
		if Contains(online, ct.exclude) {
			return false, "", nil
		}

		connect = Remove(connect, ct.exclude)
		disconnect = Remove(disconnect, ct.exclude)
	}

	var lines []string
	if len(connect) > 0 {
		lines = append(lines, fmt.Sprintf("%s connected.", FormatPlayerList(connect)))
	}
	if len(disconnect) > 0 {
		lines = append(lines, fmt.Sprintf("%s disconnected.", FormatPlayerList(disconnect)))
	}

	if len(lines) > 0 {
		return true, strings.Join(lines, "\n"), nil
	}

	return false, "", nil
}
