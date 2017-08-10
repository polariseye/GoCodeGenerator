package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type CloseCommand struct {
}

func (this *CloseCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	sessionObj.Close()
	return nil
}

func (this *CloseCommand) FlaglyDesc() string {
	return "关闭当前连接，关闭后需要open"
}

func (this *CloseCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("close",
			getHelpPcItem(),
		),
	}
}
