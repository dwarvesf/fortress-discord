package icy

import (
	"fmt"
	"strings"
	"time"
)

func formatTimeAgo(t time.Time) string {
	duration := time.Since(t)

	if duration.Seconds() < 60 {
		return fmt.Sprintf("%.0fs ago", duration.Seconds())
	} else if duration.Minutes() < 60 {
		return fmt.Sprintf("%.0fm ago", duration.Minutes())
	} else if duration.Hours() < 24 {
		return fmt.Sprintf("%.0fh ago", duration.Hours())
	} else if duration.Hours() < 720 {
		return fmt.Sprintf("%.0fd ago", duration.Hours()/24)
	} else {
		return fmt.Sprintf("%.0fM ago", duration.Hours()/720)
	}
}

// calculatePadding calculates the padding required for each column
func calculatePadding(data [][]string) []int {
	numCols := len(data[0])
	padLengths := make([]int, numCols)

	for _, row := range data {
		for i, col := range row {
			colLength := len(col)
			if colLength > padLengths[i] {
				padLengths[i] = colLength
			}
		}
	}

	return padLengths
}

// formatTable formats the data into a table with aligned columns
func formatTable(data [][]string, msgs []string) string {
	if len(data) != len(msgs) {
		return ""
	}

	padLengths := calculatePadding(data)
	var formattedRows []string

	for row, rowData := range data {
		var formattedCols []string
		for col, colData := range rowData {
			padding := strings.Repeat(" ", padLengths[col]-len(colData))

			switch col {
			case 2:
				formattedCols = append(formattedCols, fmt.Sprintf("%s`%s%s`", discordEmojiIceCube, colData, padding))
			case 3:
				formattedCols = append(formattedCols, fmt.Sprintf("%s`%s%s`", discordEmojiDollar, colData, padding))
			default:
				formattedCols = append(formattedCols, fmt.Sprintf("`%s%s`", colData, padding))
			}
		}
		formattedRows = append(formattedRows, discordEmojiSmallBlueDiamond+" "+strings.Join(formattedCols, " | "))

		msg := msgs[row]
		formattedRows = append(formattedRows, fmt.Sprintf("  \u21B3 %s", msg))
	}

	return strings.Join(formattedRows, "\n")
}
