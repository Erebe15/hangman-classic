package hangman

func DoesContain(list []string, n string) bool {
	for _, l := range list {
		if l == n {
			return true
		}
	}
	return false
}
