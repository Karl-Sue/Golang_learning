package parsinglogfiles

import (
    "regexp";
    "fmt"
)

func IsValidLine(text string) bool {
	strings := []string{
        `^\[TRC\]`,
        `^\[DBG\]`,
        `^\[INF\]`,
        `^\[WRN\]`,
        `^\[ERR\]`,
        `^\[FTL\]`,
    }
    for _,v := range strings {
        re := regexp.MustCompile(v)
        if re.MatchString(text) {
            return true
        }
    }
    return false
}

func SplitLogLine(text string) []string {
	pattern := `\<[*-=~]*\>`
    re := regexp.MustCompile(pattern)
    result := re.Split(text, -1)
    return result
}

func CountQuotedPasswords(lines []string) int {
	pattern := `"[^"]*(?i)password[^"]*"`
    re := regexp.MustCompile(pattern)
    result := 0
    for _,v := range lines {
        if re.MatchString(v) {
            result++
        }
    }
    return result
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
    re := regexp.MustCompile(pattern)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	user := `User\s+(\w+)`
    re := regexp.MustCompile(user)
    result := []string{}
    for _, v := range lines {
        matches := re.FindStringSubmatch(v) 
        if matches != nil {
            username := matches[1]
			newLine := fmt.Sprintf("[USR] %s %s", username, v)
			result = append(result, newLine)
        } else {
            result = append(result, v)
        }
    }
    return result
}
