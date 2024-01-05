package permutil

import "strings"

func CheckWhitelistChannels(whiteListedChannelsString string, channelId string) bool {
	whiteListedChannels := strings.Split(whiteListedChannelsString, ",")

	for _, id := range whiteListedChannels {
		if channelId == strings.TrimSpace(id) {
			return true
		}
	}
	return false
}
