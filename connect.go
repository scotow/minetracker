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

func OnlinePlayers(hostname string, port int, password string) ([]string, error) {
	args := []string{"-c", "-H", hostname, "-P", strconv.Itoa(port)}
	if password != "" {
		args = append(args, "-p", password)
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