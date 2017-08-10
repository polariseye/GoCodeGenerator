package templateUtil

import (
	"GoCodeGenerator/src/config"
	"GoCodeGenerator/src/dbHelper"
	"strings"
)

// 前缀处理
// item:配置项
// val:待处理的值
// 返回值:
// string:处理后的结果
func filterPrefix(item *config.FixionItem, val string) string {
	switch strings.ToLower(item.MethodType) {
	case "remove":
		return strings.TrimPrefix(val, item.Value)
	case "replace":
		return item.TargetValue + strings.TrimPrefix(val, item.Value)
	}

	return val
}

// 后缀处理
// item:配置项
// val:待处理的值
// 返回值:
// string:处理后的结果
func filterStuffix(item *config.FixionItem, val string) string {
	switch strings.ToLower(item.MethodType) {
	case "remove":
		return strings.TrimSuffix(val, item.Value)
	case "replace":
		return strings.TrimSuffix(val, item.Value) + item.TargetValue
	}

	return val
}

// 获取额外的命名空间导入项
// typeList:映射项集合
// column:待处理的列
// 返回值:
// string:结果字符串
func getExtImport(typeList []*config.MapItem, column *dbHelper.DbColumn) string {
	colType := strings.ToLower(column.DbColumnType)
	for _, item := range typeList {
		if colType == strings.ToLower(item.OldValue) {
			return item.Import
		}
	}

	colType = strings.ToLower(column.DbDataType)
	for _, item := range typeList {
		if colType == strings.ToLower(item.OldValue) {
			return item.Import
		}
	}

	return ""
}
