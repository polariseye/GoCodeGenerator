package readlineExtend

import (
	"bytes"

	"github.com/chzyer/readline"
)

// 递归自动补全结构
type RecursionCompleter struct {
	// 用于递归补全的原始对象
	source readline.PrefixCompleterInterface
}

// 信息打印
func (this *RecursionCompleter) Print(prefix string, level int, buf *bytes.Buffer) {
	this.source.Print(prefix, level, buf)
}

// 进行补全调用
func (this *RecursionCompleter) Do(line []rune, pos int) (newLine [][]rune, length int) {
	return this.source.Do(line, pos)
}

// 获取当前结构的补全名
func (this *RecursionCompleter) GetName() []rune {
	return this.source.GetName()
}

// 获取当前节点的子节点
func (this *RecursionCompleter) GetChildren() []readline.PrefixCompleterInterface {
	result := []readline.PrefixCompleterInterface{this}
	tmpChildrenList := this.source.GetChildren()
	if tmpChildrenList != nil {
		result = append(result, tmpChildrenList...)
	}

	return result
}

// 设置子节点
func (this *RecursionCompleter) SetChildren(children []readline.PrefixCompleterInterface) {
	this.source.SetChildren(children)
}

// 是否是动态的自动补全
func (this *RecursionCompleter) IsDynamic() bool {
	val, ok := this.source.(readline.DynamicPrefixCompleterInterface)
	if !ok {
		return false
	}

	return val.IsDynamic()
}

// 获取动态自动补全的名
func (this *RecursionCompleter) GetDynamicNames(line []rune) [][]rune {
	val, ok := this.source.(readline.DynamicPrefixCompleterInterface)
	if !ok {
		return [][]rune{this.GetName()}
	}

	return val.GetDynamicNames(line)
}

// 以静态补全对象作为递归补全
func PcItem(name string, pc ...readline.PrefixCompleterInterface) readline.PrefixCompleterInterface {
	tmpItem := readline.PcItem(name, pc...)
	return &RecursionCompleter{
		source: tmpItem,
	}
}

// 以动态补全对象作为递归补全
func PcItemDynamic(callback readline.DynamicCompleteFunc, pc ...readline.PrefixCompleterInterface) readline.PrefixCompleterInterface {
	tmpItem := readline.PcItemDynamic(callback, pc...)
	return &RecursionCompleter{
		source: tmpItem,
	}
}
