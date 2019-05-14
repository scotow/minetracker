package tracker

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/scotow/skyblocktracker/misc"
	. "github.com/scotow/skyblocktracker/notifier"
)

var (
	ErrInvalidOutput = errors.New("output doesn't match expected format")
	ErrCountMismatch = errors.New("number of player parsed is not matching")
)

func NewConnectionTracker(exclude string) *ConnectionTracker {
	ct := new(ConnectionTracker)
	ct.exclude = exclude
	return ct
}

type ConnectionTracker struct {
	last    []string
	exclude string
}

func (ct *ConnectionTracker) Command() string {
	return "list"
}

func (ct *ConnectionTracker) Track(result string, notifier Notifier) error {
	online, err := parseOnlinePlayers(result)
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

// TODO: Move this helper func to misc.
func parseOnlinePlayers(data string) ([]string, error) {
	fields := strings.Split(data, ":")
	if len(fields) != 2 {
		return nil, ErrInvalidOutput
	}

	fields[1] = strings.TrimSpace(fields[1])

	var count int
	n, err := fmt.Sscanf(fields[0], "There are %d of a max %d players online", &count, new(int))
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, ErrInvalidOutput
	}

	if fields[1] == "" {
		return []string{}, nil
	}

	players := strings.Split(fields[1], ", ")
	if len(players) != count {
		return nil, ErrCountMismatch
	}

	return players, nil
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
