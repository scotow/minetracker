package skyblocktracker

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var (
	ErrInvalidOutput = errors.New("output doesn't match expected format")
	ErrCountMismatch = errors.New("number of player parsed is not matching")
)

func OnlinePlayers(cred Credentials) ([]string, error) {
	args := []string{"-c", "-H", cred.Hostname, "-P", strconv.Itoa(cred.Port)}
	if cred.Password != "" {
		args = append(args, "-p", cred.Password)
	}
	args = append(args, "list")

	data, err := exec.Command("mcrcon", args...).Output()
	if _, ok := err.(*exec.ExitError); !ok {
		return nil, err
	}

	fields := strings.Split(string(data), ": ")
	if len(fields) != 2 {
		return nil, ErrInvalidOutput
	}

	fields[1] = strings.TrimSpace(fields[1])

	var count int
	n, err := fmt.Sscanf(fields[0], "There are %d of a max %d players online", &count, new(int))
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, ErrInvalidOutput
	}

	players := strings.Split(fields[1], ",")
	if len(players) != count {
		return nil, ErrCountMismatch
	}

	return players, nil
}
