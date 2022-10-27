package hangman

func NicerTypo(s string) string {
	var res string
	for _, l := range s {
		res = res + " " + string(l)
	}
	return res
}
