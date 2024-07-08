package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Viewer interface {
	Reddit(original *model.DiscordMessage, subreddit string, popular, emerging []reddit.Post) error
	Help(message *model.DiscordMessage) error
}
