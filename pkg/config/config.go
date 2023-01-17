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
	Discord   Discord

	Endpoint Endpoint
}

type Endpoint struct {
	Fortress string
}

type Discord struct {
	SecretToken string
	Prefix      string
}

type ApiServer struct {
	Port           string
	AllowedOrigins string
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
			Port: v.GetString("PORT"),
		},
		Endpoint: Endpoint{
			Fortress: v.GetString("FORTRESS_ENDPOINT"),
		},
		Discord: Discord{
			SecretToken: v.GetString("DISCORD_SECRET_TOKEN"),
			Prefix:      v.GetString("DISCORD_PREFIX"),
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
