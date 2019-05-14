package tracker

import (
	"fmt"
	"strings"
	"time"

	. "github.com/scotow/skyblocktracker/misc"
	. "github.com/scotow/skyblocktracker/notifier"
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

func (ct *ConnectionTracker) Track(result string, notifier Notifier) error {
	online, err := ParseOnlinePlayers(result)
	if err != nil {
		return err
	}

	newConnect := FindNew(ct.last, online)
	newDisconnect := FindNew(online, ct.last)

	ct.last = online

	if len(newConnect)+len(newDisconnect) > 0 {
		return ct.excludeAndNotify(online, newConnect, newDisconnect, notifier)
	}

	return nil
}

func (ct *ConnectionTracker) excludeAndNotify(online, connect, disconnect []string, notifier Notifier) error {
	if ct.exclude != "" {
		if Contains(online, ct.exclude) {
			return nil
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
		return notifier.Notify(strings.Join(lines, "\n"))
	}

	return nil
}
