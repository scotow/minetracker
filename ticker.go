package skyblocktracker

import (
	"time"

	. "github.com/scotow/skyblocktracker/notifier"
	. "github.com/scotow/skyblocktracker/tracker"
)

type Ticker struct {
	runner   Runner
	tracker  Tracker
	notifier Notifier

	interval time.Duration
	stop     chan struct{}
	report   chan<- error
}

func NewTicker(runner Runner, tracker Tracker, notifier Notifier, interval time.Duration) *Ticker {
	t := new(Ticker)
	t.runner = runner
	t.tracker = tracker
	t.notifier = notifier
	t.interval = interval
	return t
}

func (t *Ticker) Start(report chan<- error) {
	t.stop = make(chan struct{})
	t.report = report

	go func() {
		ticker := time.NewTicker(t.interval)

		for {
			select {
			case <-ticker.C:
				err := t.tick()
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
	}()
}

func (t *Ticker) Stop() {
	t.stop <- struct{}{}
}

func (t *Ticker) tick() error {
	result, err := t.runner.RunCommand(t.tracker.Command())
	if err != nil {
		return err
	}

	err = t.tracker.Track(result, t.notifier)
	if err != nil {
		return err
	}

	return nil
}
