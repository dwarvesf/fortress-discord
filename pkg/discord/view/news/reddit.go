package news

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func (v view) Reddit(original *model.DiscordMessage, subreddit string, popular, emerging []reddit.Post) error {
	content := make([]string, 0)

	title := fmt.Sprintf("**<:pepe_ping:1028964391690965012> Reddit %s BUZZ!! <:pepe_ping:1028964391690965012>**", strings.ToUpper(subreddit))

	if len(popular) > 0 {
		content = append(content, "<a:badge5:1131851001117294672> **POPULAR**")
		for _, post := range popular {
			content = append(content, fmt.Sprintf("[[%v](%s)] %s", post.NumberOfComments, "https://www.reddit.com"+post.Permalink, post.Title))
		}
		content = append(content, "[See more...](https://www.reddit.com/r/golang/rising/)")
	}

	// Separate popular and emerging
	content = append(content, "")

	if len(emerging) > 0 {
		content = append(content, "<a:arrow_up_animated:1131789319644921936> **EMERGING**")
		for _, post := range emerging {
			content = append(content, fmt.Sprintf("[[%s](%s)] %s", timeAgo(post.Created.Time), "https://www.reddit.com"+post.Permalink, post.Title))
		}
		content = append(content, "[See more...](https://www.reddit.com/r/golang/new/)")
	}

	msg := &discordgo.MessageEmbed{
		Title:       title,
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
