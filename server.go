package minetracker

import (
	. "github.com/scotow/minetracker/notifier"
	. "github.com/scotow/minetracker/runner"
	. "github.com/scotow/minetracker/tracker"
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
