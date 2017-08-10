package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type HelpCommand struct {
}

func (this *HelpCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	println(h.GetRoot().Usage(""))

	return nil
}

func (this *HelpCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("help"),
	}
}

func (this *HelpCommand) FlaglyDesc() string {
	return "帮助"
}
