package minetracker

import (
	"time"

	. "github.com/scotow/minetracker/notifier"
	. "github.com/scotow/minetracker/runner"
	. "github.com/scotow/minetracker/tracker"
)

type Ticker struct {
	runner   Runner
	tracker  Tracker
	notifier Notifier
	stop     chan struct{}
}

func NewTicker(runner Runner, tracker Tracker, notifier Notifier) *Ticker {
	t := new(Ticker)
	t.runner = runner
	t.tracker = tracker
	t.notifier = notifier
	return t
}

func (t *Ticker) Start(report chan<- error) {
	t.stop = make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(t.tracker.Wait()):
				err := t.tick()
				if err != nil {
					if report != nil {
						report <- err
					}
					return
				}
			case <-t.stop:
				return
			}
		}
	}()
}

func (t *Ticker) Stop() {
	t.stop <- struct{}{}
}

func (t *Ticker) tick() error {
	result, err := t.runner.Run(t.tracker.Command())
	if err != nil {
		return err
	}

	should, result, err := t.tracker.Track(result)
	if err != nil {
		return err
	}

	if should {
		err = t.notifier.Notify(result)
		if err != nil {
			return err
		}
	}

	return nil
}
