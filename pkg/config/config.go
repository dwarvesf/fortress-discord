package config

import (
	"github.com/spf13/viper"
)

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

type Config struct {
	Debug bool
	Env   string

	ApiServer ApiServer
	OpenAI    OpenAI
	Discord   Discord
	Reddit    Reddit

	Endpoint Endpoint

	DiscordResearchTopic DiscordResearchTopic
}

type Reddit struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

type OpenAI struct {
	APIKey string
}

type Endpoint struct {
	Fortress string
	Mochi    string
}

type Discord struct {
	SecretToken         string
	Prefix              string
	ID                  DiscordIds
	WhiteListedChannels string
}

type DiscordIds struct {
	FortressBot       string
	DwarvesGuild      string
	RepostDoneChannel string
}

type ApiServer struct {
	APIKey         string
	Port           string
	AllowedOrigins string
}

type DiscordResearchTopic struct {
	Size string
}

type ENV interface {
	GetBool(string) bool
	GetString(string) string
}

func Generate(v ENV) *Config {
	return &Config{
		Debug: v.GetBool("DEBUG"),
		Env:   v.GetString("ENV"),
		ApiServer: ApiServer{
			Port:   v.GetString("PORT"),
			APIKey: v.GetString("FORTRESS_API_KEY"),
		},
		Endpoint: Endpoint{
			Fortress: v.GetString("FORTRESS_ENDPOINT"),
			Mochi:    v.GetString("MOCHI_ENDPOINT"),
		},
		OpenAI: OpenAI{
			APIKey: v.GetString("OPENAI_API_KEY"),
		},
		Discord: Discord{
			SecretToken: v.GetString("DISCORD_SECRET_TOKEN"),
			Prefix:      v.GetString("DISCORD_PREFIX"),
			ID: DiscordIds{
				FortressBot:       v.GetString("DISCORD_ID_FORTRESS_BOT"),
				DwarvesGuild:      v.GetString("DISCORD_ID_DWARVES_GUILD"),
				RepostDoneChannel: v.GetString("DISCORD_ID_REPOST_DONE"),
			},
			WhiteListedChannels: v.GetString("DISCORD_WHITELISTED_CHANNELS"),
		},
		Reddit: Reddit{
			ClientID:     v.GetString("REDDIT_CLIENT_ID"),
			ClientSecret: v.GetString("REDDIT_CLIENT_SECRET"),
			Username:     v.GetString("REDDIT_USERNAME"),
			Password:     v.GetString("REDDIT_PASSWORD"),
		},
		DiscordResearchTopic: DiscordResearchTopic{
			Size: v.GetString("RESEARCH_TOPIC_SIZE"),
		},
	}
}

func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) *Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "local")
	v.SetDefault("DISCORD_PREFIX", "?")
	v.SetDefault("RESEARCH_TOPIC_SIZE", "5")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return Generate(v)
}

func LoadTestConfig() Config {
	return Config{
		Debug: true,
	}
}
