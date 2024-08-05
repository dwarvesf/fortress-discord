package adapter

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/mochi"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/openai"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/tono"
)

type IAdapter interface {
	Fortress() fortress.FortressAdapter
	Mochi() mochi.MochiAdapter
	OpenAI() openai.OpenAIAdapter
	Tono() tono.TonoAdapter
}
