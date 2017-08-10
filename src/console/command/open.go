package command

import (
	"GoCodeGenerator/src/session"

	"errors"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type OpenCommand struct {
	Ip         string `name:"h" desc:"数据库IP"`
	Port       int    `name:"p" desc:"数据库端口"`
	UserName   string `name:"usr" desc:"用户名"`
	Pwd        string `name:"pwd" desc:"用户密码"`
	DefaultDb  string `name:"defaultdb" desc:"连接的默认数据库"`
	ConfigName string `name:"c" desc:"已有的连接配置"`
}

func (this *OpenCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	if len(this.ConfigName) > 0 {
		return sessionObj.OpenFromConfig(this.ConfigName)
	}

	return sessionObj.Open(this.UserName, this.Pwd, this.Ip, this.Port, this.DefaultDb)
}

// 参数检查
func (this *OpenCommand) FlaglyVerify() error {
	if len(this.ConfigName) > 0 {
		return nil
	}

	if len(this.Ip) <= 0 {
		return errors.New("error:ip is empty")
	}
	if this.Port <= 0 {
		return errors.New("error:port is empty")
	}
	if len(this.UserName) <= 0 {
		return errors.New("error:username is empty")
	}

	return nil
}

func (this *OpenCommand) FlaglyDesc() string {
	return "打开数据库连接"
}

func (this *OpenCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{readline.PcItem("open",
		getHelpPcItem(),
		readline.PcItem("-h"),
		readline.PcItem("-p"),
		readline.PcItem("-pwd"),
		readline.PcItem("-usr"),
		readline.PcItem("-defaultdb"),
		readline.PcItem("-c",
			readline.PcItemDynamic(func(val string) []string {
				return sessionObj.QueryConnectionConfig("")
			}),
		)),
	}
}
