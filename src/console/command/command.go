package command

import (
	"GoCodeGenerator/src/session"
	"GoCodeGenerator/src/util/uprint"
	"reflect"

	"github.com/alecthomas/kingpin"
	"github.com/chzyer/readline"
)

var cmdCreateFun map[string]func(_sessionObj *session.BuildSession, app *kingpin.Application) PrefixCompleter

type PrefixCompleter interface {
	GetPrefixCompleter(sessionObj *session.BuildSession) []readline.PrefixCompleterInterface
}

// 获取指定对象的所有字段的自动补全数据
func GetFieldCompleterList(obj interface{}, sessionObj *session.BuildSession) []readline.PrefixCompleterInterface {
	val := reflect.ValueOf(obj).Elem()
	result := make([]readline.PrefixCompleterInterface, 0)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanInterface() == false {
			continue
		}

		completer := field.Interface().(PrefixCompleter)

		// 获取字段的自动补全数据
		childCompleterList := completer.GetPrefixCompleter(sessionObj)
		if childCompleterList != nil {
			result = append(result, childCompleterList...)
		}
	}

	return result
}

func printList(valList []string) {
	uprint.Printf("\r\n")
	lines := ""
	for index, db := range valList {
		lines += db + "    "
		if index > 0 && index%5 == 0 {
			lines += "\r\n"
		}
	}
	lines += "\r\n"
	uprint.Print(lines)
}

func getHelpPcItem() readline.PrefixCompleterInterface {
	return readline.PcItem("-help")
}
