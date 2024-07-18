package news

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// Render renders news message to Discord, using for any platform. Will migrate reddit later
func (v view) Render(original *model.DiscordMessage, platform, topic string, popular, emerging []model.News) error {
	content := make([]string, 0)

	title := fmt.Sprintf("**<:pepe_ping:1028964391690965012> %s %s BUZZ!! <:pepe_ping:1028964391690965012>**", strings.ToUpper(platform), strings.ToUpper(topic))

	if len(popular) > 0 {
		content = append(content, "<a:badge5:1131851001117294672> **POPULAR**")
		for _, post := range popular {
			content = append(content, fmt.Sprintf("[[%v](%s)] %s", post.Popularity, post.URL, post.Title))
		}
		content = append(content, seeMore(platform, topic))
	}

	// Separate popular and emerging
	content = append(content, "")

	if len(emerging) > 0 {
		content = append(content, "<a:arrow_up_animated:1131789319644921936> **EMERGING**")
		for _, post := range emerging {
			content = append(content, fmt.Sprintf("[[%s](%s)] %s", timeAgo(post.CreatedAt), post.URL, post.Title))
		}
		content = append(content, seeMore(platform, topic))
	}

	if len(popular) == 0 && len(emerging) == 0 {
		content = append(content, "Oops! No posts found. Make sure your platform and topic are valid!")
	}

	msg := &discordgo.MessageEmbed{
		Title:       title,
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
