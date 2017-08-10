package main

import (
	"os"

	"github.com/chzyer/readline"
)

func main() {
	Start()
}

func Start() {
	var err error
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "tmp/prompt.txt",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    readline.NewPrefixCompleter(getCompleter()),

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer rl.Close()
	line, _ := rl.Readline()
	println(line)
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func getCompleter() readline.PrefixCompleterInterface {
	return PcItem("test")
}
