package probleme

func IsPalindrome(word string) bool {
	return reverse(word) == word

}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func Palindromes(words []string) (l []string) {
	var list []string

	for _, w := range words {
		if IsPalindrome(w) {
			list = append(list, w)
		}
	}

	return list
}
