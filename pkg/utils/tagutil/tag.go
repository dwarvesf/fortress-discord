package tagutil

import "fmt"

// FormatRole return discord format for tagging role
func FormatRole(role string) string {
	return fmt.Sprintf("<@&%s>", role)
}
