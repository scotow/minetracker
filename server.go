package skyblocktracker

import (
	"os/exec"
	"strconv"
	"time"

	. "github.com/scotow/skyblocktracker/notifier"
	. "github.com/scotow/skyblocktracker/tracker"
)

type Runner interface {
	RunCommand(string) (string, error)
}

func NewServer(hostname string, port int, password string, report chan<- error) *Server {
	s := new(Server)
	s.hostname = hostname
	s.port = port
	s.password = password
	s.report = report
	return s
}

type Server struct {
	hostname string
	port     int
	password string
	report   chan<- error
}

func (s *Server) RunCommand(cmd string) (string, error) {
	args := []string{"-c", "-H", s.hostname, "-P", strconv.Itoa(s.port)}
	if s.password != "" {
		args = append(args, "-p", s.password)
	}
	args = append(args, "list")

	data, err := exec.Command("mcrcon", args...).Output()
	if _, ok := err.(*exec.ExitError); !ok {
		return "", err
	}

	return string(data), nil
}

func (s *Server) Add(tracker Tracker, notifier Notifier, interval time.Duration) *Ticker {
	ticker := NewTicker(s, tracker, notifier, interval)
	ticker.Start(s.report)

	return ticker
}
