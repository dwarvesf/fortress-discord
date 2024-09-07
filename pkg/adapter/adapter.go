package adapter

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter/dify"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/mochi"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/openai"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/tono"
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
	Tono     tono.TonoAdapter
	Dify     dify.DifyAdapter
}

func New(cfg *config.Config, l logger.Logger) IAdapter {
	return &Adapter{
		subAdapter: subAdapter{
			Fortress: fortress.New(cfg.Endpoint.Fortress, cfg.ApiServer.APIKey),
			Mochi:    mochi.New(cfg.Endpoint.Mochi),
			OpenAI:   openai.New(cfg.OpenAI.APIKey),
			Tono:     tono.New(cfg),
			Dify:     dify.New(cfg.Dify.BaseURL, cfg.Dify.SummarizerAppToken, cfg.Dify.ProcessAIAppToken),
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

func (a *Adapter) Tono() tono.TonoAdapter {
	return a.subAdapter.Tono
}

func (a *Adapter) Dify() dify.DifyAdapter {
	return a.subAdapter.Dify
}
