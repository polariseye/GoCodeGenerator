package templateUtil

import (
	"GoCodeGenerator/src/config"
	"GoCodeGenerator/src/dbHelper"
	"reflect"
	"strings"
	"time"
)

type TemplateFun struct {
}

var templateFunData map[string]interface{} = make(map[string]interface{})

// 分隔符
const con_Seprator = "_"

func init() {
	funObj := new(TemplateFun)

	val := reflect.ValueOf(funObj)
	tp := reflect.TypeOf(funObj)

	// 把函数名和方法保存到集合中
	for i := 0; i < tp.NumMethod(); i++ {
		methodName := tp.Method(i).Name
		templateFunData[methodName] = val.MethodByName(methodName).Interface()
	}
}

// 首字母大写
// str:待处理的字符串
// 返回值:
// string:结果
func (this *TemplateFun) FirstCharUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

// 首字母小写
// str:待处理的字符串
// 返回值:
// string:结果
func (this *TemplateFun) FirstCharLower(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

// 获取实体名
// tableName:表名
// 返回值:
// string:实体名
func (this *TemplateFun) EntityName(tableName string) string {
	if len(tableName) <= 0 {
		return ""
	}

	// 前缀处理
	var result = tableName
	for _, item := range config.GetTemplateConfig().TablePrefix {
		if strings.HasPrefix(result, item.Value) {
			result = filterPrefix(item, result)
			break
		}
	}

	// 后缀处理
	for _, item := range config.GetTemplateConfig().TableStuffix {
		if strings.HasSuffix(result, item.Value) {
			result = filterStuffix(item, result)
			break
		}
	}

	wordsList := strings.Split(result, con_Seprator)
	for index, item := range wordsList {
		wordsList[index] = strings.ToUpper(item[0:1]) + item[1:]
	}

	return strings.Join(wordsList, "")
}

// 数据库类型转换为对应的数据类型
// dbDataType:数据库的数据类型
// dbColumnType:数据库的完整声明
// 返回值:
// string:go数据类型
func (this *TemplateFun) FieldType(column *dbHelper.DbColumn) string {
	typeList := config.GetTemplateConfig().DataTypeMapTable

	colType := strings.ToLower(column.DbColumnType)
	for _, item := range typeList {
		if colType == strings.ToLower(item.OldValue) {
			return item.TargetValue
		}
	}

	colType = strings.ToLower(column.DbDataType)
	for _, item := range typeList {
		if colType == strings.ToLower(item.OldValue) {
			return item.TargetValue
		}
	}

	return colType
}

// 获取首字母大写的字段名
// name:字段名
// 返回值:
// string:字段名
func (this *TemplateFun) FieldName(name string) string {
	colMapList := config.GetTemplateConfig().ColumnMapTable
	for _, item := range colMapList {
		if item.OldValue == name {
			name = item.TargetValue
		}
	}

	return this.FirstCharUpper(name)
}

// 获取额外的导入信息
// column:待处理的列项
// 返回值:
// string:导入字符串
func (this *TemplateFun) GetExtImport(tableItem *dbHelper.DbTable) []string {
	typeList := config.GetTemplateConfig().DataTypeMapTable

	result := make([]string, 0)
	for _, colItem := range tableItem.Columns {
		// 获取导入项
		importItem := getExtImport(typeList, colItem)
		if len(importItem) <= 0 {
			continue
		}

		// 如果已经存在，则跳过
		ifExist := false
		for _, item := range result {
			if item == importItem {
				ifExist = true
			}
		}
		if ifExist {
			continue
		}

		result = append(result, importItem)
	}

	return result
}

// 把字符串拼接为一个字符串
// a:字符串列表
// sep:间隔符
// 返回值:
// string:拼接后的字符串
func (this *TemplateFun) Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

// 给所有字段加上前缀
// columns:列集合
// Postfix:要加的前缀
// sep:字段间的间隔
// 返回值:
// 列的字符串
func (this *TemplateFun) ColumnWithPostfix(columns []string, Postfix, sep string) string {
	result := make([]string, 0, len(columns))

	for _, t := range columns {
		result = append(result, t+Postfix)
	}

	return strings.Join(result, sep)
}

// 字符串重复
func (this *TemplateFun) Repeat(val string, num int) string {
	return strings.Repeat(val, num)
}

// 获取列表中字符串的最大长度
// vals:待处理的字符串列表
// 返回值:
// int:最大长度
func (this *TemplateFun) MaxLen(vals []string) int {
	maxLen := 0
	for _, item := range vals {
		if maxLen < len(item) {
			maxLen = len(item)
		}
	}

	return maxLen + 1
}

// 获取对齐后的字符串
func (this *TemplateFun) Assign(baseVal string, repeatVal string, maxLen int) string {
	if len(baseVal) >= maxLen {
		return baseVal
	}

	return baseVal + this.Repeat(repeatVal, maxLen-len(baseVal))
}

// 获取主键列集合
// table:数据表对象
func (this *TemplateFun) GetColumns(table *dbHelper.DbTable, ifPrimary bool) []*dbHelper.DbColumn {
	result := make([]*dbHelper.DbColumn, 0)

	for _, item := range table.Columns {
		if ifPrimary == item.IsPrimaryKey {
			result = append(result, item)
		}
	}

	return result
}

// 按照名字排除要获取的列项
// table:表对象
// exclueColNames:要排除的列名集合
// 返回值:
// []*dbHelper.DbColumn:列集合
func (this *TemplateFun) GetColumnsExcloud(table *dbHelper.DbTable, exclueColNames ...string) []*dbHelper.DbColumn {
	result := make([]*dbHelper.DbColumn, 0)

	for _, item := range table.Columns {
		ifExist := false
		for _, colNameItem := range exclueColNames {
			if colNameItem == item.ColumnName {
				ifExist = true
			}
		}

		if ifExist == false {
			result = append(result, item)
		}
	}

	return result
}

// 获取当前时间
func (this *TemplateFun) Now() time.Time {
	return time.Now()
}

// 时间格式化输出
func (this *TemplateFun) TimeFormat(targetTime time.Time, format string) string {
	if len(format) <= 0 {
		format = "2006-01-02 15:04:05"
	}

	return targetTime.Local().Format(format)
}

// 获取用于模版处理的函数列表
// 返回值:
// map[string]interface{}:函数映射列表
func GetTemplateFunData() map[string]interface{} {
	result := make(map[string]interface{})

	// 返回函数列表
	for key, val := range templateFunData {
		result[key] = val
	}

	return result
}

// 默认的对象
func GetTemplateFunObj() *TemplateFun {
	return &TemplateFun{}
}
