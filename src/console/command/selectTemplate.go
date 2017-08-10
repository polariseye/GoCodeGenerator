package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type SelectTemplateCommand struct {
	TargetTemplate string `name:"template" type:"[0]" desc:"要使用的模板"`
}

func (this *SelectTemplateCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	return sessionObj.SelectTemplateGroup(this.TargetTemplate)
}

func (this *SelectTemplateCommand) FlaglyDesc() string {
	return "选择要使用的模版"
}

func (this *SelectTemplateCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("selecttemplate",
			readline.PcItemDynamic(func(val string) []string {
				return sessionObj.QueryTemplateGroup("")
			}),
		),
	}
}
