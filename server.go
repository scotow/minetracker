package minetracker

import (
	. "github.com/scotow/minetracker/notifier"
	. "github.com/scotow/minetracker/runner"
	. "github.com/scotow/minetracker/tracker"
)

// Create a Server.
// runner is the runner that will be used by the Tickers.
// report is the channel used to report errors reported by the Trackers.
func NewServer(runner Runner, report chan<- error) *Server {
	s := new(Server)
	s.runner = runner
	s.report = report
	return s
}

// Server is a object used to group multiple Trackers on the same Runner and help creating Ticker.
type Server struct {
	runner Runner
	report chan<- error
}

// Link the Tracker and the Notifier.
func (s *Server) Add(tracker Tracker, notifier Notifier) *Ticker {
	ticker := NewTicker(s.runner, tracker, notifier)
	ticker.Start(s.report)
	return ticker
}
