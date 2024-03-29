package misc

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidOutput = errors.New("output doesn't match expected format")
	ErrCountMismatch = errors.New("number of player parsed is not matching")
)

var (
	PlayerJoinString     = ", "
	PlayerLastJoinString = " and "
)

// Compare the two slices and return the elements that wasn't in the first slice.
func FindNew(old, new []string) []string {
	diff := make([]string, 0)
	for _, s := range new {
		if !Contains(old, s) {
			diff = append(diff, s)
		}
	}
	return diff
}

// A simple function that indicates if a string is present in a slice of strings.
func Contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// Find and remove a specific string in a slice of strings.
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// A function that concatenate players string with a giver string, and an other string for final join.
// Set the PlayerJoinString and PlayerLastJoinString exposed variables for customization.
func FormatPlayerList(players []string) string {
	switch len(players) {
	case 0:
		return ""
	case 1:
		return players[0]
	default:
		return fmt.Sprintf("%s%s%s", strings.Join(players[:len(players)-1], PlayerJoinString), PlayerLastJoinString, players[len(players)-1])
	}
}

// Parse a string of the "There are %d of a max 3 players online: a, b, c" format and return the list of online players.
func ParseOnlinePlayers(data string) ([]string, error) {
	fields := strings.Split(data, ":")
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

	if fields[1] == "" {
		return []string{}, nil
	}

	players := strings.Split(fields[1], ", ")
	if len(players) != count {
		return nil, ErrCountMismatch
	}

	return players, nil
}
