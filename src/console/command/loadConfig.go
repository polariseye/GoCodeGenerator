package command

import (
	"GoCodeGenerator/src/session"

	"errors"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type LoadConfigCommand struct {
	FileName string `name:"fl" type:"[0]" desc:"配置文件名"`
}

func (this *LoadConfigCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	if len(this.FileName) <= 0 {
		return errors.New("error:please input config file name")
	}

	sessionObj.LoadConfig(this.FileName)
	return nil
}

func (this *LoadConfigCommand) FlaglyDesc() string {
	return "加载配置文件"
}

func (this *LoadConfigCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("loadconfig",
			getHelpPcItem(),
			readline.PcItemDynamic(func(pre string) []string {
				fls, _ := sessionObj.GetConfigFileNames()
				return fls
			}),
		),
	}
}
