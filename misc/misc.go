package misc

import (
	"fmt"
	"strings"
)

func FindNew(old, new []string) []string {
	diff := make([]string, 0)
	for _, s := range new {
		if !Contains(old, s) {
			diff = append(diff, s)
		}
	}
	return diff
}

func Contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func FormatPlayerList(players []string) string {
	switch len(players) {
	case 0:
		return ""
	case 1:
		return players[0]
	default:
		return fmt.Sprintf("%s and %s", strings.Join(players[0:len(players)-1], ", "), players[len(players)-1])
	}
}
