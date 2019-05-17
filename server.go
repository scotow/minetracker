package skyblocktracker

import (
	. "github.com/scotow/skyblocktracker/notifier"
	. "github.com/scotow/skyblocktracker/runner"
	. "github.com/scotow/skyblocktracker/tracker"
)

func NewServer(runner Runner, report chan<- error) *Server {
	s := new(Server)
	s.runner = runner
	s.report = report
	return s
}

type Server struct {
	runner Runner
	report chan<- error
}

func (s *Server) Add(tracker Tracker, notifier Notifier) *Ticker {
	ticker := NewTicker(s.runner, tracker, notifier)
	ticker.Start(s.report)
	return ticker
}
