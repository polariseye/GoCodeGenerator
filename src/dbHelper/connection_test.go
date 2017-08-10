package dbHelper

import (
	"testing"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func TestDb(context *testing.T) {
	connectionObj, errMsg := NewConnection("root", "1111", "127.0.0.1", 3306)
	if errMsg != nil {
		context.Error(errMsg.Error())

		return
	}

	var dbInfo *DbInfo
	if dbInfo, errMsg = connectionObj.LoadDb("trunk"); errMsg != nil {
		context.Error(errMsg.Error())

		return
	}

	fmt.Printf("表数量：%v\r\n", len(dbInfo.tables))
	for _, table := range dbInfo.tables {
		fmt.Printf("\r\n开始输出表:%v comment:%v", table.TableName, table.Comment)
		for _, col := range table.Columns {
			fmt.Printf("\r\n	ColumnName:%v DbDataType=%v IsAutoIncrement=%v IsNullAble=%v IsPrimaryKey=%v IsUnsigned=%v Len=%v Scale=%v Comment=%v DbColumnType=%v",
				col.ColumnName,
				col.DbDataType,
				col.IsAutoIncrement,
				col.IsNullAble,
				col.IsPrimaryKey,
				col.IsUnsigned,
				col.Len,
				col.Scale,
				col.Comment,
				col.DbColumnType)
		}
	}
}

func TestLoadTable(context *testing.T) {
	connectionObj, errMsg := NewConnection("root", "11111", "127.0.0.1", 3306)
	if errMsg != nil {
		context.Error(errMsg.Error())

		return
	}

	var tableItem *DbTable
	if tableItem, errMsg = connectionObj.LoadTable("trunk", "table1"); errMsg != nil {
		context.Error(errMsg.Error())

		return
	}

	for _, col := range tableItem.Columns {
		fmt.Printf("\r\n	ColumnName:%v DbDataType=%v IsAutoIncrement=%v IsNullAble=%v IsPrimaryKey=%v IsUnsigned=%v Len=%v Scale=%v Comment=%v DbColumnType=%v",
			col.ColumnName,
			col.DbDataType,
			col.IsAutoIncrement,
			col.IsNullAble,
			col.IsPrimaryKey,
			col.IsUnsigned,
			col.Len,
			col.Scale,
			col.Comment,
			col.DbColumnType)
	}
}
