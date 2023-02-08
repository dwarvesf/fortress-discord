package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

func Test_shouldParseMessage(t *testing.T) {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	l := logger.NewLogrusLogger()

	// dummy discord session
	ses := discordgo.Session{}

	d := New(&ses, cfg, l, nil, nil)

	type tc struct {
		name string
		msg  *discordgo.MessageCreate
		want bool
	}

	cases := []tc{{
		name: "valid should parse",
		msg: &discordgo.MessageCreate{
			Message: &discordgo.Message{
				Content: "?earn",
				Author: &discordgo.User{
					ID: "123",
				},
				Member: &discordgo.Member{
					Roles: []string{"123"},
				},
			},
		},
		want: true,
	},
		{
			name: "invalid prefix",
			msg: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "!earn",
					Author: &discordgo.User{
						ID: "123",
					},
				},
			},
			want: false,
		},
		{
			name: "invalid message",
			msg: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: "",
					Author: &discordgo.User{
						ID: "123",
					},
				},
			},
			want: false,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			got := d.shouldParseMessage(c.msg)

			if got != c.want {
				t.Errorf("want %v, got %v", c.want, got)
			}
		})
	}
}

func Test_parseMessage(t *testing.T) {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	l := logger.NewLogrusLogger()

	// dummy discord session
	ses := discordgo.Session{}

	d := New(&ses, cfg, l, nil, nil)

	type tc struct {
		name string
		msg  *discordgo.MessageCreate
	}

	cases := []tc{{
		name: "valid should parse",
		msg: &discordgo.MessageCreate{
			Message: &discordgo.Message{
				Content:   "?earn list",
				ChannelID: "123456",
				GuildID:   "987654321",
				Author: &discordgo.User{
					ID: "123",
				},
				Member: &discordgo.Member{
					Roles: []string{"123"},
				},
			},
		},
	}}

	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			got := d.parseMessage(c.msg)

			if got.RawContent != c.msg.Content {
				t.Errorf("want %v, got %v", c.msg.Content, got.RawContent)
			}

			if got.ChannelId != c.msg.ChannelID {
				t.Errorf("want %v, got %v", c.msg.ChannelID, got.ChannelId)
			}

			if got.GuildId != c.msg.GuildID {
				t.Errorf("want %v, got %v", c.msg.GuildID, got.GuildId)
			}

			if got.Author.ID != c.msg.Author.ID {
				t.Errorf("want %v, got %v", c.msg.Author.ID, got.Author.ID)
			}

			if got.ContentArgs[0] != "earn" {
				t.Errorf("want %v, got %v", "earn", got.ContentArgs[0])
			}

			if got.ContentArgs[1] != "list" {
				t.Errorf("want %v, got %v", "list", got.ContentArgs[0])
			}
		})
	}
}
