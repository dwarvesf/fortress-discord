package model

type TemplateType string

const (
	TemplateSummary   TemplateType = "summary"
	TemplateKeyPoints TemplateType = "key_points"
	TemplateNuggets   TemplateType = "nuggets"
	TemplateFacts     TemplateType = "facts"
)

func IsValidTemplateType(t string) bool {
	switch TemplateType(t) {
	case TemplateSummary, TemplateKeyPoints, TemplateNuggets, TemplateFacts:
		return true
	default:
		return false
	}
}