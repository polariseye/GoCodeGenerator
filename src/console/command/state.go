package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type StateCommand struct {
}

func (this *StateCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	println(sessionObj.StateVal())
	return nil
}

func (this *StateCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("state"),
	}
}

func (this *StateCommand) FlaglyDesc() string {
	return "查看当前状态"
}
