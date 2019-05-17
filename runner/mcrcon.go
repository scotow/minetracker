package runner

import (
	"os/exec"
	"strconv"
	"sync"
)

var (
	mcrconPath = "mcrcon"
)

func SetMcrconPath(path string) {
	mcrconPath = path
}

func NewMcrcon(cred Credentials) *McrconServer {
	ms := new(McrconServer)
	ms.cred = cred
	return ms
}

type McrconServer struct {
	cred Credentials
	lock sync.Mutex
}

func (ms *McrconServer) Run(cmd string) (string, error) {
	ms.lock.Lock()
	defer ms.lock.Unlock()

	args := []string{"-c", "-H", ms.cred.Hostname, "-P", strconv.Itoa(ms.cred.Port)}
	if ms.cred.Password != "" {
		args = append(args, "-p", ms.cred.Password)
	}
	args = append(args, cmd)

	data, err := exec.Command(mcrconPath, args...).Output()
	if _, ok := err.(*exec.ExitError); !ok {
		return "", err
	}

	return string(data), nil
}
