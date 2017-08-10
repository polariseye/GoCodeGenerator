package command

import (
	"GoCodeGenerator/src/session"
	"os"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type QuitCommand struct {
}

func (this *QuitCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	os.Exit(0)
	return nil
}
func (this *QuitCommand) FlaglyDesc() string {
	return "关闭程序ctrl+c"
}

func (this *QuitCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("quit"),
	}
}
