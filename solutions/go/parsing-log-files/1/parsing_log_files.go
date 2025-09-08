package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]\s.+$`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`^.*["]{1}.*((?i)password).*["]{1}.*$`)

	matches := 0

	for _, l := range lines {
		if found := re.FindAllString(l, -1); found != nil {
			matches += len(found)
		}
	}

	return matches
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line)[0-9]*`)

	modifiedText := re.ReplaceAllString(text, "")

	return modifiedText
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	modifiedLines := make([]string, 0, len(lines))

	for _, l := range lines {
		if matches := re.FindStringSubmatch(l); matches != nil {
			username := matches[1]
			l = fmt.Sprintf("[USR] %s %s", username, l)
		}
		modifiedLines = append(modifiedLines, l)
	}
	return modifiedLines
}
