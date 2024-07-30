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

	if len(emerging) > 0 {
		emerging = emerging[0:5]
		content = append(content, "<a:arrow_up_animated:1131789319644921936> **EMERGING**")
		for _, post := range emerging {
			postContent := fmt.Sprintf("[[%s](%s)] **%s** \n", timeAgo(post.CreatedAt), post.URL, post.Title)
			postContent += v.getAdditionContent(post)
			content = append(content, postContent)
		}
		content = append(content, seeMore(platform, topic))
	}

	// Separate popular and emerging
	content = append(content, "")

	if len(popular) > 0 {
		popular = popular[0:5]
		content = append(content, "<a:badge5:1131851001117294672> **POPULAR**")
		index := 1
		for _, post := range popular {
			postContent := fmt.Sprintf("[[%v](%s)] **%s** \n", index, post.URL, post.Title)
			postContent += v.getAdditionContent(post)
			content = append(content, postContent)

			index++
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

func (v view) getAdditionContent(news model.News) string {
	if len(news.Tags) > 0 {
		return fmt.Sprintf("`∟ Score: %d • Comments: %d • Tags: %s`\n", news.Popularity, news.CommentCount, strings.Join(news.Tags, ", "))
	}
	return fmt.Sprintf("`∟ Score: %d • Comments: %d`\n", news.Popularity, news.CommentCount)
}
