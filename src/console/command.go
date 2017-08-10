package console

import (
	"GoCodeGenerator/src/console/command"
	"GoCodeGenerator/src/session"

	"github.com/chzyer/readline"
)

type AppCommand struct {
	Help           *command.HelpCommand           `name:"help" flagly:"handler" `
	Quit           *command.QuitCommand           `name:"quit" flagly:"handler" `
	LoadConfig     *command.LoadConfigCommand     `name:"loadconfig" flagly:"handler" `
	Open           *command.OpenCommand           `name:"open" flagly:"handler" `
	Reload         *command.ReloadCommand         `name:"reload" flagly:"handler" `
	Close          *command.CloseCommand          `name:"close" flagly:"handler" `
	Db             *command.DbCommand             `name:"db" flagly:"handler" `
	SelectDb       *command.SelectDbCommand       `name:"selectdb" flagly:"handler" `
	Table          *command.TableCommand          `name:"table" flagly:"handler" `
	SelectTemplate *command.SelectTemplateCommand `name:"selecttemplate" flagly:"handler" `
	Build          *command.BuildCommand          `name:"build" flagly:"handler" `
	State          *command.StateCommand          `name:"state" flagly:"handler" `
}

func (this *AppCommand) getCompleter(sessionObj *session.BuildSession) *readline.PrefixCompleter {
	result := command.GetFieldCompleterList(this, sessionObj)

	return readline.NewPrefixCompleter(result...)
}

func getAppCommandObj() *AppCommand {
	return &AppCommand{}
}
