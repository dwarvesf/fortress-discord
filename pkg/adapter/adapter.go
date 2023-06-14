package adapter

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/mochi"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/openai"
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type Adapter struct {
	subAdapter
}

type subAdapter struct {
	Fortress fortress.FortressAdapter
	Mochi    mochi.MochiAdapter
	OpenAI   openai.OpenAIAdapter
}

func New(cfg *config.Config, l logger.Logger) IAdapter {
	return &Adapter{
		subAdapter: subAdapter{
			Fortress: fortress.New(cfg.Endpoint.Fortress, cfg.ApiServer.APIKey),
			Mochi:    mochi.New(cfg.Endpoint.Mochi),
			OpenAI:   openai.New(cfg.OpenAI.APIKey),
		},
	}
}

func (a *Adapter) Fortress() fortress.FortressAdapter {
	return a.subAdapter.Fortress
}

func (a *Adapter) Mochi() mochi.MochiAdapter {
	return a.subAdapter.Mochi
}

// OpenAI implements IAdapter.
func (a *Adapter) OpenAI() openai.OpenAIAdapter {
	return a.subAdapter.OpenAI
}
