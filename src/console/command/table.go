package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type TableCommand struct {
	Table string `name:"db" type:"[0]" desc:"目标数据库"`
}

func (this *TableCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	tbLst, err := sessionObj.QueryTable(this.Table)
	if err != nil {
		return err
	}

	if len(tbLst) <= 0 {
		println("table no found")
		return nil
	}

	printList(tbLst)

	return nil
}

func (this *TableCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("table",
			getHelpPcItem(),
			getEnumTablePcItem(sessionObj),
		),
	}
}

func (this *TableCommand) FlaglyDesc() string {
	return "列取数据表"
}

func getEnumTablePcItem(sessionObj *session.BuildSession) readline.PrefixCompleterInterface {
	return readline.PcItemDynamic(func(val string) []string {
		result, _ := sessionObj.QueryTable("")
		return result
	})
}
