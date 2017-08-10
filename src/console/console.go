package console

import (
	"GoCodeGenerator/src/session"
	"GoCodeGenerator/src/util/uprint"
	"os"

	"strings"

	"github.com/chzyer/flagly"
	"github.com/chzyer/readline"
	"github.com/google/shlex"
)

func Start() {
	cmdObj := getAppCommandObj()
	sessionObj := session.NewSession()
	fset, err := flagly.Compile("", cmdObj)
	fset.Context(sessionObj) //// 关联到命令中
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "tmp/prompt.txt",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    cmdObj.getCompleter(sessionObj),

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer rl.Close()
	uprint.SetOutput(rl)

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) <= 0 {
			continue
		}

		command, err := shlex.Split(line)
		if err != nil {
			println("error: " + err.Error())
			continue
		}

		if err := fset.Run(command); err != nil {
			println(err.Error())
		}
	}
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
