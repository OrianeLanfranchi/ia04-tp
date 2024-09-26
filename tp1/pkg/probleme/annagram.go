package probleme

import (
	"sort"
	"strings"
)

func Footprint(s string) (footprint string) {
	var sSort = strings.Split(s, "")
	sort.Strings(sSort)
	return strings.Join(sSort, "")
}

func Anagrams(words []string) (anagrams map[string][]string) {
	dict := make(map[string][]string)

	for _, e := range words {
		dict[Footprint(e)] = append(dict[Footprint(e)], e)
	}

	return dict
}
