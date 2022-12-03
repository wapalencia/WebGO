package internal
func IsPalindrome(w string) bool {

	i := 0
	for _, item := range w {
		i++
		rev := item
		ini := w[len(w)-i]
		if string(rev) != string(ini) {
			println("el for", string(rev), "vs cadena", string(ini))
			return false
		}
	}
	return true
}