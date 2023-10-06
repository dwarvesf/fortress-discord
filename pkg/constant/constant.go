package constant

import "os"

// tech radar
const (
	TechRadarRingTrial  string = "Trial"
	TechRadarRingAdopt         = "Adopt"
	TechRadarRingAssess        = "Assess"
	TechRadarRingHold          = "Hold"
	TechRadarAll               = ""
)

const (
	DiscordBraineryBot     = "1020554094705909820"
	DiscordBraineryChannel = "955015316293972048"

	DiscordPlayGroundBraineryBot     = "1119172751891120208"
	DiscordPlayGroundBraineryChannel = "1119171172198797393"
)

const (
	RegexPatternDiscordChannelID = `<#(\d+)>`
	RegexPatternDiscordID        = `<@(\d+)>`
	RegexPatternEmail            = `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+\b`
	RegexPatternIcyReward        = ` (\d+)`
	RegexPatternNumber           = `\d{18,}`
	RegexPatternUrl              = `((?:https?://)[^\s]+)`
	RegexPatternGithub           = `gh:(\w+)`
	RegexPatternDescription      = `d:"(.*?)"`
)

const (
	DefaultBraineryReward = "0"
)

var mapEmoji = map[string]string{
	"ARROW_DOWN_ANIMATED": "<a:arrow_down_animated:1131789144759214171>",
	"ARROW_UP_ANIMATED":   "<a:arrow_up_animated:1131789319644921936>",
	"BADGE1":              "<a:badge1:1131850989062852638>",
	"BADGE2":              "<a:badge2:1131850991663337554>",
	"BADGE3":              "<a:badge3:1131850996159610930>",
	"BADGE5":              "<a:badge5:1131851001117294672>",
	"LOG_CHANNEL":         "<:log_channel:1131863319377100841>",
	"STAR_ANIMATED":       "<a:star_animated:1131862886592024586>",
	"INCREASING_ANIMATED": "<a:increasing_animated:1131862879319097394>",
	"CLOCK_NEW":           "<:clock_new:1131863089185292428>",
}

func GetEmoji(emoji string) string {
	if os.Getenv("ENV") != "prod" {
		return mapEmojiDev[emoji]
	}
	return mapEmoji[emoji]
}

var mapEmojiDev = map[string]string{
	"ARROW_DOWN_ANIMATED": "<a:arrow_up_animated:1131317348670902292>",
	"ARROW_UP_ANIMATED":   "<a:arrow_down_animated:1131317344774397992>",
	"BADGE1":              "<a:badge1:1133460615684440167>",
	"BADGE2":              "<a:badge2:1133460619253796914>",
	"BADGE3":              "<a:badge3:1133460622365958304>",
	"BADGE5":              "<a:badge5:1133460625784320021>",
	"LOG_CHANNEL":         "<:logchannel:1133460455906627614>",
	"STAR_ANIMATED":       "<a:star_animated:1133460443550195832>",
	"INCREASING_ANIMATED": "<a:increasing_animated:1133460451091550289>",
	"CLOCK_NEW":           "<:clock:1133460445257281658>",
}

var DwarvesRole = map[string]bool{
	"moderator":  true,
	"dwarf":      true,
	"booster":    true,
	"apprentice": true,
	"crafter":    true,
	"specialist": true,
	"principal":  true,
	"peeps":      true,
	"learning":   true,
	"engagement": true,
	"delivery":   true,
	"labs":       true,
	"baby dwarf": true,
	"ladies":     true,
	"sers":       true,
	"consultant": true,
	"chad":       true,
}
