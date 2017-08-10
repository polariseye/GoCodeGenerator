package dbHelper

// 数据库表信息
type DbTable struct {
	TableName string      // 表名
	Comment   string      // 描述
	Columns   []*DbColumn // 列集合
}

// 获取所有列的名称
// 返回值:
// []string:所有列名
func (this *DbTable) ALlColumnNames() []string {
	colNames := make([]string, len(this.Columns))
	for _, colItem := range this.Columns {
		colNames = append(colNames, colItem.ColumnName)
	}

	return colNames
}

// 获取所有列的类型
// 返回值:
// []string:所有列字段类型
func (this *DbTable) AllColumnTypes() []string {
	colTypes := make([]string, len(this.Columns))
	for _, colItem := range this.Columns {
		colTypes = append(colTypes, colItem.DbDataType)
	}

	return colTypes
}

// 创建数据库表对象
// tableName:表名
// comment:表注释
// columns:列集合
func newDbTable(tableName string, comment string, columns []*DbColumn) *DbTable {
	return &DbTable{
		TableName: tableName,
		Comment:   comment,
		Columns:   columns,
	}
}
