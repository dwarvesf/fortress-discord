package stringutils

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"regexp"
	"sort"
	"strings"
	"time"

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
