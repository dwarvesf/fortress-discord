package news

import (
	"fmt"
	"time"
)

// timeAgo converts a time to a string representation of how long ago it was.
func timeAgo(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	if duration.Hours() >= 24 {
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d day%s ago", days, pluralize(days))
	} else if duration.Hours() >= 1 {
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hour%s ago", hours, pluralize(hours))
	} else {
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d minute%s ago", minutes, pluralize(minutes))
	}
}

// Helper function to add 's' for pluralization.
func pluralize(count int) string {
	if count != 1 {
		return "s"
	}
	return ""
}

func seeMore(platform, topic string) string {
	switch platform {
	case "reddit":
		return fmt.Sprintf("[See more...](https://www.reddit.com/r/%s/new/)\n", topic)
	case "lobsters":
		return fmt.Sprintf("[See more...](https://lobste.rs/t/%s)\n", topic)
	default:
		return ""
	}
}
