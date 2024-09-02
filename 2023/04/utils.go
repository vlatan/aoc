package scratchcards

import "strings"

func removeEmpty(s []string) []string {
	result := []string{}
	for _, v := range s {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func getSets(line string) ([]string, []string) {
	_, sets, _ := strings.Cut(line, ": ")
	w, s, _ := strings.Cut(sets, " | ")
	winning := strings.Split(w, " ")
	scratched := strings.Split(s, " ")
	return removeEmpty(winning), removeEmpty(scratched)
}
