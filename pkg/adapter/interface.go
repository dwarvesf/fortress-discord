package adapter

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter/dify"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/icy"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/ir"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/mochi"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/n8n"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/openai"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/tono"
)

type IAdapter interface {
	Fortress() fortress.FortressAdapter
	Mochi() mochi.MochiAdapter
	OpenAI() openai.OpenAIAdapter
	Tono() tono.TonoAdapter
	Dify() dify.DifyAdapter
	IR() ir.IRAdapter
	Icy() icy.IcyAdapter
	N8n() n8n.N8nAdapter
}
