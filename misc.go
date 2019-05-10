package skyblocktracker

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
	for _, i := range ss {
		if i == s {
			return true
		}
	}
	return false
}
