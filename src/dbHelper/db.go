package dbHelper

import (
	"database/sql"
	"fmt"
	"strconv"
)

// 数据库信息
type DbInfo struct {
	connectionObj *Connection         //// 数据库连接诶对象
	dbName        string              //// 数据库名
	tables        map[string]*DbTable //// 表集合
}

// 加载一个数据库的所有表
// dbName:数据库名
// 返回值:
// error:错误信息
func (this *DbInfo) Reload() error {
	result, errMsg := getAllDbTable(this.connectionObj.GetConnection(), this.dbName)
	if errMsg != nil {
		return errMsg
	}

	tables := make(map[string]*DbTable)
	for _, item := range result {
		tables[item.TableName] = item
	}

	this.tables = tables

	return nil
}

// 获取所有表名
// 返回值:
// []string:表名集合
func (this *DbInfo) GetTableNames() []string {
	tableNames := make([]string, 0, len(this.tables))

	for key, _ := range this.tables {
		tableNames = append(tableNames, key)
	}

	return tableNames
}

// 获取指定表对象
// tableName:表名
// *DbTable:表对象
// bool:是否存在此表对象
func (this *DbInfo) GetTable(tableName string) (*DbTable, bool) {
	result, isExist := this.tables[tableName]

	return result, isExist
}

// 新建一个数据库对象
// _connectionObj:连接对象
// _dbName:数据库名
// 返回值:
// *DbInfo:数据库对象
// error:错误信息
func newDbInfo(_connectionObj *Connection, _dbName string) (*DbInfo, error) {
	dbObj := &DbInfo{
		connectionObj: _connectionObj,
		dbName:        _dbName,
	}

	errMsg := dbObj.Reload()
	if errMsg != nil {
		return nil, errMsg
	}

	return dbObj, nil
}

// 读取表的所有列数据
// db:数据库连接对象
// dbName:数据库名
// tableName:表名
// 返回值:
// []*DbColumn:列集合
// error:错误信息
func getALLColumns(db *sql.DB, dbName string, tableName string) ([]*DbColumn, error) {
	executeSql := `SELECT 
COLUMN_NAME,IS_NULLABLE,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,NUMERIC_PRECISION,NUMERIC_SCALE
,COLUMN_KEY,EXTRA,COLUMN_COMMENT,COLUMN_TYPE
 from COLUMNS
where TABLE_NAME=? and table_schema = ?;`
	rows, errMsg := db.Query(executeSql, tableName, dbName)
	if errMsg != nil {
		return nil, fmt.Errorf("执行获取数据库表字段异常，错误信息:%v", errMsg.Error())
	}

	var (
		columnName   []byte // 预防为null导致异常
		isNullAble   []byte
		dbDataType   []byte
		dbColumnType []byte
		stringLen    []byte
		numLen       []byte
		scaleLen     []byte
		columnKey    []byte
		extra        []byte
		comment      []byte
	)

	result := make([]*DbColumn, 0)
	for rows.Next() {
		errMsg = rows.Scan(&columnName, &isNullAble, &dbDataType, &stringLen,
			&numLen, &scaleLen, &columnKey, &extra, &comment, &dbColumnType)

		if errMsg != nil {
			return nil, errMsg
		}

		result = append(result, newDbColumn(string(columnName), string(isNullAble),
			string(dbDataType), string(dbColumnType), bytesToInt(stringLen),
			bytesToInt(numLen), bytesToInt(scaleLen), string(columnKey), string(extra), string(comment)))
	}

	return result, nil
}

// 获取所有表信息
// db:数据库连接对象
// dbName:数据库名
// 返回值:
// []*DbTable:表集合
// error:错误信息
func getAllDbTable(db *sql.DB, dbName string) ([]*DbTable, error) {
	executeSql := "SELECT TABLE_NAME,TABLE_COMMENT FROM tables WHERE table_schema=?"
	rows, errMsg := db.Query(executeSql, dbName)
	if errMsg != nil {
		return nil, fmt.Errorf("获取数据库所有表失败:%v", errMsg.Error())
	}

	var (
		tableName string
		comment   string
		columns   []*DbColumn
	)

	result := make([]*DbTable, 0)
	for rows.Next() {
		errMsg = rows.Scan(&tableName, &comment)
		if errMsg != nil {
			return nil, errMsg
		}

		columns, errMsg = getALLColumns(db, dbName, tableName)
		if errMsg != nil {
			return nil, errMsg
		}

		result = append(result, newDbTable(tableName, comment, columns))
	}

	return result, nil
}

// byte转换成整型
// data:字节数组
// 返回值:
// int:结果值
func bytesToInt(data []byte) int {
	val, _ := strconv.ParseInt(string(data), 10, 32)

	return int(val)

	/*	b_buf := bytes.NewBuffer(data)
		var x int
		fmt.Println("字节数据:", data)
		binary.Write(b_buf, binary.LittleEndian, x)

		return x
	*/
}
