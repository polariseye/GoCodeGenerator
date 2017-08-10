package command

import (
	"GoCodeGenerator/src/session"

	"errors"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type SelectDbCommand struct {
	Db string `name:"db" type:"[0]" desc:"目标数据库"`
}

func (this *SelectDbCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	return sessionObj.SelectDb(this.Db)
}

// 参数检查
func (this *SelectDbCommand) FlaglyVerify() error {
	if len(this.Db) < 0 {
		return errors.New("error: db is empty")
	}

	return nil
}

func (this *SelectDbCommand) FlaglyDesc() string {
	return "选择要操作的数据库"
}

func (this *SelectDbCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("selectdb",
			getHelpPcItem(),
			getEnumDbPcItem(sessionObj),
		),
	}
}
