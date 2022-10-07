package hangman

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			return true
		}
		if s[i] > 64 && s[i] < 91 {
			return true
		}
	}
	return false
}
