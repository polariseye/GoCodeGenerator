package dbHelper

import (
	"strings"
)

// 数据库列信息
type DbColumn struct {
	ColumnName      string // 字段名
	DbDataType      string // 数据库类型
	Len             int    // 长度
	Scale           int    // 精度
	IsNullAble      bool   // 是否能为null
	IsAutoIncrement bool   // 是否是自增
	IsPrimaryKey    bool   // 是否是主键
	IsUnsigned      bool   // 是否是无符号类型
	Comment         string // 备注信息
	DbColumnType    string // 数据库中的列类型声明字符串
}

// 创建新行
// columnName:列名
// isNullAble:是否可以为空 true false
// dbDataType:数据库中的数据类型
// dbColumnType:数据库列的类型信息
// stringLen:数据库中字符串长度
// numLen:数据库中数值类型的长度
// scaleLen:数值类型的小数位数
// columnKey:列的索引信息，主键或外键的信息
// extra:额外信息 是否是自增
// comment:备注信息
func newDbColumn(columnName string, isNullAble string, dbDataType string, dbColumnType string, stringLen int,
	numLen int, scaleLen int, columnKey string, extra string, comment string) *DbColumn {
	dataLen := stringLen
	if numLen > 0 {
		dataLen = numLen
	}

	result := &DbColumn{
		ColumnName:      columnName,
		DbDataType:      strings.ToLower(dbDataType),
		Len:             dataLen,
		Scale:           scaleLen,
		IsNullAble:      strings.ToLower(isNullAble) == "yes",
		IsAutoIncrement: strings.Contains(strings.ToLower(extra), "auto_increment"),
		IsUnsigned:      strings.Contains(strings.ToLower(dbColumnType), "unsigned"),
		Comment:         comment,
		IsPrimaryKey:    strings.Contains(strings.ToLower(columnKey), "pri"),
		DbColumnType:    dbColumnType,
	}

	return result
}
