package message

import (
	"strconv"
	"strings"
)

func Count(s string) int {
	return len(s)
}

func CountCharactersInt(s int) int {
	return Count(strconv.Itoa(s))
}

func CountSubstrings(subject, substring string) int {
	restString := subject
	substringLength := len(substring)
	substringCount := 0

	if (subject == "") || (substring == "") {
		return 0
	}

	for {
		matchIndex := strings.Index(restString, substring)

		if matchIndex == -1 {
			break
		}

		substringCount++
		if matchIndex+substringLength >= len(restString) {
			break
		}
		restString = restString[matchIndex+substringLength:]
	}

	return substringCount
}
