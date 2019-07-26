package minetracker

import (
	"time"

	. "github.com/scotow/minetracker/notifier"
	. "github.com/scotow/minetracker/runner"
	. "github.com/scotow/minetracker/tracker"
)

// Create a Ticker.
// runner is the Runner used to run the Tracker's commands.
// tracker is Tracker that provide the command and the interval between two runs.
// notifier is the Notifier used to send the notifications provided by the Tracker.
func NewTicker(runner Runner, tracker Tracker, notifier Notifier) *Ticker {
	t := new(Ticker)
	t.runner = runner
	t.tracker = tracker
	t.notifier = notifier
	return t
}

// A Ticker used a Runner to run command provided by Tracker and notify using a Notifier if the Tracker found something.
type Ticker struct {
	runner   Runner
	tracker  Tracker
	notifier Notifier
	stop     chan struct{}
}

// Start the Ticker on its own goroutine.
// report is the channel used to report error if one occurred.
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

// Stop the Ticker and its associated goroutine.
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
