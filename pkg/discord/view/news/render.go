package news

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// Render renders news message to Discord, using for any platform.
func (v view) Render(original *model.DiscordMessage, platform, topic string, posts []model.News) error {
	title := fmt.Sprintf("<:pepe_ping:1028964391690965012> %s %s BUZZ!! <:pepe_ping:1028964391690965012>", strings.ToUpper(platform), strings.ToUpper(topic))

	if len(posts) == 0 {
		msg := &discordgo.MessageEmbed{
			Title:       title,
			Description: "Oops! No posts found. Make sure your platform and topic are valid!",
			Color:       0xFF0000, // Red color for error message
		}
		return base.SendEmbededMessage(v.ses, original, msg)
	}

	// Limit to 10 posts to keep the message concise
	maxPosts := 10
	if len(posts) > maxPosts {
		posts = posts[:maxPosts]
	}

	var description strings.Builder
	for i, post := range posts {
		description.WriteString(v.formatPostContent(post, i+1))
		description.WriteString("\n\n")
	}

	description.WriteString(seeMore(platform, topic))

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: description.String(),
		Color:       0x00BFFF, // DeepSkyBlue color
	}

	return base.SendEmbededMessage(v.ses, original, embed)
}

func (v view) formatPostContent(post model.News, index int) string {
	var title string
	if post.URL != "" {
		title = fmt.Sprintf("[[%d]](%s) **%s**", index, post.URL, post.Title)
	} else {
		title = fmt.Sprintf("[%d] %s", index, post.Title)
	}

	additionalContent := v.getAdditionalContent(post)
	return fmt.Sprintf("%s\n%s", title, additionalContent)
}

func (v view) getAdditionalContent(news model.News) string {

	score := fmt.Sprintf("📊 Score: %d", news.Popularity)
	comments := fmt.Sprintf("💬 Comments: %d", news.CommentCount)
	posted := fmt.Sprintf("🕒 Posted: %s", timeAgo(news.CreatedAt))

	content := fmt.Sprintf("%s | %s | %s", score, comments, posted)

	if len(news.Tags) > 0 {
		tags := fmt.Sprintf("🏷️ Tags: *%s*", strings.Join(news.Tags, ", "))
		content += fmt.Sprintf(" | %s", tags)
	}

	return content
}
