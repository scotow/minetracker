package skyblocktracker

import (
	"time"
)

type TrackChange func(online, connect, disconnect []string)

type Tracker struct {
	cred     Credentials
	last     []string
	change   TrackChange
	interval time.Duration

	stop   chan struct{}
	report chan<- error
}

func NewTracker(cred Credentials, change TrackChange, interval time.Duration) *Tracker {
	t := new(Tracker)
	t.cred = cred
	t.change = change
	t.interval = interval

	return t
}

func (t *Tracker) Start(report chan<- error) error {
	last, err := OnlinePlayers(t.cred)
	if err != nil {
		return err
	}

	t.last = last
	t.stop = make(chan struct{})
	t.report = report

	go t.startInterval()
	return nil
}

func (t *Tracker) startInterval() {
	ticker := time.NewTicker(t.interval)

	for {
		select {
		case <-ticker.C:
			err := t.updateAndNotify()
			if err != nil {
				ticker.Stop()
				t.report <- err
				return
			}
		case <-t.stop:
			ticker.Stop()
			return
		}
	}
}

func (t *Tracker) updateAndNotify() error {
	online, err := OnlinePlayers(t.cred)
	if err != nil {
		return err
	}

	newConnect := make([]string, 0)
	for _, player := range online {
		if t.isNew(player) {
			newConnect = append(newConnect, player)
		}
	}

	newDisconnect := make([]string, 0)
	for _, player := range t.last {
		if t.wasConnected(player) {
			newConnect = append(newConnect, player)
		}
	}

	if len(newConnect)+len(newDisconnect) > 0 {
		t.change(online, newConnect, newDisconnect)
	}

	t.last = online
	return nil
}

func (t *Tracker) isNew(player string) bool {
	for _, p := range t.last {
		if p == player {
			return false
		}
	}
	return true
}

func (t *Tracker) wasConnected(player string) bool {
	for _, p := range t.last {
		if p == player {
			return true
		}
	}
	return false
}
