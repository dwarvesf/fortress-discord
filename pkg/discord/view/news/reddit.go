package news

import (
	"fmt"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func (v view) Reddit(original *model.DiscordMessage, subreddit string, popular, emerging []reddit.Post) error {
	content := make([]string, 0)

	title := fmt.Sprintf("**<:pepe_ping:1028964391690965012> Reddit %s BUZZ!! <:pepe_ping:1028964391690965012>**", strings.ToUpper(subreddit))

	if len(emerging) > 0 {
		content = append(content, "<a:arrow_up_animated:1131789319644921936> **EMERGING**")
		for _, post := range emerging {
			escapedTitle := stringutils.EscapeMarkdown(post.Title)
			postContent := fmt.Sprintf("[[%s](%s)] **%s** \n", timeAgo(post.Created.Time), "https://www.reddit.com"+post.Permalink, escapedTitle)
			postContent += fmt.Sprintf("`∟ Score: %d • Comments: %d • Upvote ratio: %.2f`\n", post.Score, post.NumberOfComments, post.UpvoteRatio)
			content = append(content, postContent)
		}
		content = append(content, "[See more...](https://www.reddit.com/r/golang/new/)")
	}

	// Separate popular and emerging
	content = append(content, "")

	if len(popular) > 0 {
		index := 1
		content = append(content, "<a:badge5:1131851001117294672> **POPULAR**")
		for _, post := range popular {
			escapedTitle := stringutils.EscapeMarkdown(post.Title)
			postContent := fmt.Sprintf("[[%v](%s)] **%s** \n", index, "https://www.reddit.com"+post.Permalink, escapedTitle)
			postContent += fmt.Sprintf("`∟ Score: %d • Comments: %d • Upvote ratio: %.2f`\n", post.Score, post.NumberOfComments, post.UpvoteRatio)
			content = append(content, postContent)

			index++
		}
		content = append(content, "[See more...](https://www.reddit.com/r/golang/rising/)")
	}

	msg := &discordgo.MessageEmbed{
		Title:       title,
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
