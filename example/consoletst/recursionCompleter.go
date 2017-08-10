package main

import (
	"bytes"

	"github.com/chzyer/readline"
)

type RecursionCompleter struct {
	source readline.PrefixCompleterInterface
}

func (this *RecursionCompleter) Print(prefix string, level int, buf *bytes.Buffer) {
	this.source.Print(prefix, level, buf)
}
func (this *RecursionCompleter) Do(line []rune, pos int) (newLine [][]rune, length int) {
	return this.source.Do(line, pos)
}
func (this *RecursionCompleter) GetName() []rune {
	return this.source.GetName()
}
func (this *RecursionCompleter) GetChildren() []readline.PrefixCompleterInterface {
	result := []readline.PrefixCompleterInterface{this}
	tmpChildrenList := this.source.GetChildren()
	if tmpChildrenList != nil {
		result = append(result, tmpChildrenList...)
	}

	// println("进行一次调用:", len(result), "\r\n")
	return result
}
func (this *RecursionCompleter) SetChildren(children []readline.PrefixCompleterInterface) {
	this.source.SetChildren(children)
}

func (this *RecursionCompleter) IsDynamic() bool {
	val, ok := this.source.(readline.DynamicPrefixCompleterInterface)
	if !ok {
		return false
	}

	return val.IsDynamic()
}

func (this *RecursionCompleter) GetDynamicNames(line []rune) [][]rune {
	val, ok := this.source.(readline.DynamicPrefixCompleterInterface)
	if !ok {
		return [][]rune{this.GetName()}
	}

	return val.GetDynamicNames(line)
}

func PcItem(name string, pc ...readline.PrefixCompleterInterface) readline.PrefixCompleterInterface {
	tmpItem := readline.PcItem(name, pc...)
	return &RecursionCompleter{
		source: tmpItem,
	}
}

func PcItemDynamic(callback readline.DynamicCompleteFunc, pc ...readline.PrefixCompleterInterface) readline.PrefixCompleterInterface {
	tmpItem := readline.PcItemDynamic(callback, pc...)
	return &RecursionCompleter{
		source: tmpItem,
	}
}
