package runner

import (
	"os/exec"
	"strconv"
	"sync"
)

var (
	mcrconPath = "mcrcon"
)

// Change the path of the mcrcon command. (default is "mcrcon")
func SetMcrconPath(path string) {
	mcrconPath = path
}

// Create a CommandRunner using the passed Credentials.
// cred is the Credentials used to send commands.
func NewCommandRunner(cred Credentials) *CommandRunner {
	cr := new(CommandRunner)
	cr.cred = cred
	return cr
}

// The CommandRunner used the mcrcon command and the os/exec package to run the commands on the RCON server.
type CommandRunner struct {
	cred Credentials
	lock sync.Mutex
}

func (cr *CommandRunner) Run(cmd string) (string, error) {
	cr.lock.Lock()
	defer cr.lock.Unlock()

	args := []string{"-c", "-H", cr.cred.Hostname, "-P", strconv.Itoa(cr.cred.Port)}
	if cr.cred.Password != "" {
		args = append(args, "-p", cr.cred.Password)
	}
	args = append(args, cmd)

	data, err := exec.Command(mcrconPath, args...).Output()
	if _, ok := err.(*exec.ExitError); !ok {
		return "", err
	}

	return string(data), nil
}
