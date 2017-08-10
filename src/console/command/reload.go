package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type ReloadCommand struct {
}

func (this *ReloadCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	return sessionObj.Refresh()
}

func (this *ReloadCommand) FlaglyDesc() string {
	return "重新加载所有信息"
}

func (this *ReloadCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("reload",
			getHelpPcItem(),
		),
	}
}
