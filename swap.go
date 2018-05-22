package StringUtil

import "strings"

func Swap(s string) string {
	words := strings.Fields(s)
	outputWord := ""
	for i := len(words)-1; i > -1; i = i-1 {
		outputWord += words[i] + " "
	}
	return string(outputWord)
}
