package stringutils

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
)

func ExtractPattern(str string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(str, -1)
	var result []string
	for _, match := range matches {
		result = append(result, match[1])
	}

	return result
}

func ExtractEmailPattern(str string) []string {
	re := regexp.MustCompile(constant.RegexPatternEmail)
	matches := re.FindAllStringSubmatch(str, -1)
	var result []string
	for _, match := range matches {
		result = append(result, match[0])
	}

	return result
}

func ExtractNumber(str string) []string {
	re := regexp.MustCompile(constant.RegexPatternNumber)
	matches := re.FindAllStringSubmatch(str, -1)
	var result []string
	for _, match := range matches {
		result = append(result, match[0])
	}

	return result
}

func FormatString(str string) string {
	// Replace spaces with a single space
	re := regexp.MustCompile(`\s+`)
	formattedStr := re.ReplaceAllString(str, " ")

	// Remove spaces after the "#" symbol
	formattedStr = strings.ReplaceAll(formattedStr, "# ", "#")

	return formattedStr
}

func GetKeysFromMap[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ConvertToTimeAgo(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	switch {
	case duration.Seconds() < 60:
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	case duration.Minutes() < 60:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration.Hours() < 24:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration.Hours() < 48:
		return "yesterday"
	case duration.Hours() < 24*7:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	case duration.Hours() < 24*30:
		return fmt.Sprintf("%d weeks ago", int(duration.Hours()/(24*7)))
	case duration.Hours() < 24*365:
		return fmt.Sprintf("%d months ago", int(duration.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%d years ago", int(duration.Hours()/(24*365)))
	}
}

func SortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func EscapeMarkdown(s string) string {
	// Escape underscores and asterisks
	s = strings.ReplaceAll(s, "_", "\\_")
	return s
}

func ParseTimePeriod(now time.Time, timeStr string) (*time.Time, int, string, error) {
	re := regexp.MustCompile(`(?i)^(\d+)\s*([a-z]+)$`)
	matches := re.FindStringSubmatch(timeStr)
	if len(matches) != 3 {
		return nil, 0, "", errors.New("invalid duration argument format")
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, 0, "", errors.New("invalid number in duration argument")
	}

	unit := matches[2]
	var from time.Time
	switch strings.ToLower(unit) {
	case "d", "day", "days":
		from = now.AddDate(0, 0, -num)
		if num > 1 {
			return &from, num, "days", nil
		}
		return &from, num, "day", nil
	case "w", "week", "weeks":
		from = now.AddDate(0, 0, -7*num)
		if num > 1 {
			return &from, num, "weeks", nil
		}
		return &from, num, "week", nil
	case "m", "month", "months":
		from = now.AddDate(0, -num, 0)
		if num > 1 {
			return &from, num, "months", nil
		}
		return &from, num, "month", nil
	case "y", "year", "years":
		from = now.AddDate(-num, 0, 0)
		if num > 1 {
			return &from, num, "years", nil
		}
		return &from, num, "year", nil
	default:
		return nil, 0, "", errors.New("invalid time duration unit")
	}
}

// ExtractDiscordID extracts number from this format <@421992793582469130>
func ExtractDiscordID(id string) string {
	// Use a regular expression to extract the number
	re := regexp.MustCompile(`<@(\d+)>`)
	matches := re.FindStringSubmatch(id)

	// If there's a match, return the extracted number
	if len(matches) == 2 {
		return matches[1]
	}

	// If no match is found, return an empty string
	return ""
}
