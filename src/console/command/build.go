package command

import (
	"GoCodeGenerator/src/session"
	"GoCodeGenerator/src/util/readlineExtend"
	"errors"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
)

type BuildCommand struct {
	Like   bool   `name:"l" desc:"是否生成所有包含指定字符串的表" default:"false"`
	OutPut string `name:"o" desc:"输出目录" default:"resultcode/"`

	Content []string `name:"content" type:"[]" desc:"是否生成所有包含指定字符串的表"`
}

func (this *BuildCommand) FlaglyHandle(h *flagly.Handler, sessionObj *session.BuildSession) error {
	if this.Like {
		return sessionObj.BuildByLike(this.OutPut, this.Content[0])
	}

	return sessionObj.Build(this.OutPut, this.Content...)
}

func (this *BuildCommand) FlaglyDesc() string {
	return "使用模版生成代码"
}

// 参数检查
func (this *BuildCommand) FlaglyVerify() error {
	if len(this.OutPut) <= 0 {
		return errors.New("error:output path is empty")
	}
	if len(this.Content) <= 0 {
		return errors.New("error:table name is empty")
	}

	return nil
}

func (this *BuildCommand) GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("build",
			getHelpPcItem(),
			readlineExtend.PcItemDynamic(func(val string) []string {
				result, _ := sessionObj.QueryTable("")
				return result
			}),
			readline.PcItem("-l"),
		),
	}
}
