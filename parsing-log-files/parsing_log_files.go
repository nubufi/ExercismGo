package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=/-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
	re := regexp.MustCompile(pattern)

	// Replace all occurrences of the pattern with an empty string
	cleanedString := re.ReplaceAllString(text, "")

	return cleanedString
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\S+)`
	re := regexp.MustCompile(pattern)

	var result []string
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			userName := matches[1]
			taggedLine := fmt.Sprintf("[USR] %s %s", userName, line)
			result = append(result, taggedLine)
		} else {
			result = append(result, line)
		}
	}
	return result
}
