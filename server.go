package skyblocktracker

import (
	"os/exec"
	"strconv"
	"sync"

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
	lock     sync.Mutex
}

func (s *Server) RunCommand(cmd string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	args := []string{"-c", "-H", s.hostname, "-P", strconv.Itoa(s.port)}
	if s.password != "" {
		args = append(args, "-p", s.password)
	}
	args = append(args, cmd)

	data, err := exec.Command("mcrcon", args...).Output()
	if _, ok := err.(*exec.ExitError); !ok {
		return "", err
	}

	return string(data), nil
}

func (s *Server) Add(tracker Tracker, notifier Notifier) *Ticker {
	ticker := NewTicker(s, tracker, notifier)
	ticker.Start(s.report)

	return ticker
}
