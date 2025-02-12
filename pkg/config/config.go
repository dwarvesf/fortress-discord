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
	Dify      Dify
	N8n       N8n

	Endpoint Endpoint
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

type IR struct {
	Url    string
	APIKey string
}
type Endpoint struct {
	Fortress string
	Mochi    string
	Tono     string
	IR       IR
	Icy      string
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
	RandomChannel     string
	BraineryChannel   string
	DevChannel        string
}

type ApiServer struct {
	APIKey         string
	Port           string
	AllowedOrigins string
}

type Dify struct {
	BaseURL            string
	SummarizerAppToken string
	ProcessAIAppToken  string
}

type N8n struct {
	WebhookURL      string
	WebhookUsername string
	WebhookPassword string
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
			Tono:     v.GetString("TONO_ENDPOINT"),
			IR:       IR{Url: v.GetString("IR_ENDPOINT"), APIKey: v.GetString("IR_API_KEY")},
			Icy:      v.GetString("ICY_BASE_URL"),
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
				RandomChannel:     v.GetString("DISCORD_ID_RANDOM_CHANNEL"),
				BraineryChannel:   v.GetString("DISCORD_ID_BRAINERY_CHANNEL"),
				DevChannel:        v.GetString("DISCORD_ID_DEV_CHANNEL"),
			},
			WhiteListedChannels: v.GetString("DISCORD_WHITELISTED_CHANNELS"),
		},
		Reddit: Reddit{
			ClientID:     v.GetString("REDDIT_CLIENT_ID"),
			ClientSecret: v.GetString("REDDIT_CLIENT_SECRET"),
			Username:     v.GetString("REDDIT_USERNAME"),
			Password:     v.GetString("REDDIT_PASSWORD"),
		},
		Dify: Dify{
			BaseURL:            v.GetString("DIFY_BASE_URL"),
			SummarizerAppToken: v.GetString("DIFY_SUMMARIZER_APP_TOKEN"),
			ProcessAIAppToken:  v.GetString("DIFY_PROCESS_AI_APP_TOKEN"),
		},
		N8n: N8n{
			WebhookURL:      v.GetString("N8N_WEBHOOK_URL"),
			WebhookUsername: v.GetString("N8N_WEBHOOK_USERNAME"),
			WebhookPassword: v.GetString("N8N_WEBHOOK_PASSWORD"),
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
