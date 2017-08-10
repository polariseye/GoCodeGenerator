package command

import (
	"GoCodeGenerator/src/session"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type DbCommand struct {
	Db string `name:"db" type:"[0]" desc:"目标数据库"`
}

func (this *DbCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	dbLst, err := sessionObj.QueryDb(this.Db)
	if err != nil {
		return err
	}

	if len(dbLst) <= 0 {
		println("db no found")
		return nil
	}

	printList(dbLst)

	return nil
}

func (this *DbCommand) FlaglyDesc() string {
	return "列取数据库"
}

func (this *DbCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("db",
			getHelpPcItem(),
			getEnumDbPcItem(sessionObj),
		),
	}
}

func getEnumDbPcItem(sessionObj *session.BuildSession) readline.PrefixCompleterInterface {
	return readline.PcItemDynamic(func(val string) []string {
		dbList, _ := sessionObj.QueryDb("")
		return dbList
	})
}
