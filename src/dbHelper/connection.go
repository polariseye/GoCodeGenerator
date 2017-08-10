package dbHelper

import (
	"database/sql"
	"fmt"
)

// 连接对象
type Connection struct {
	connectionObj *sql.DB //// 数据库连接对象
	ip            string  //// IP
	port          int     //// 端口
	userName      string  //// 用户名
	passward      string  //// 密码
	dbName        string  //// 数据库名
}

// 获取数据库连接对象
// 返回值:
// *sql.DB:数据库连接对象
func (this *Connection) GetConnection() *sql.DB {
	return this.connectionObj
}

// 关闭数据库连接
func (this *Connection) CloseConnection() {
	if this.connectionObj != nil {
		this.connectionObj.Close()
	}
}
func (this *Connection) String() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)",
		this.userName, this.passward, this.ip, this.port)
}

// 获取所有数据库名
// 返回值:
// []string:数据库名列表
// error:错误信息
func (this *Connection) GetDataBaseNames() ([]string, error) {
	executeSql := "SELECT SCHEMA_NAME FROM SCHEMATA"
	rows, errMsg := this.connectionObj.Query(executeSql)
	if errMsg != nil {
		return nil, fmt.Errorf("获取数据库所有表失败:%v", errMsg.Error())
	}

	result := make([]string, 0)
	var dbName string
	for rows.Next() {
		errMsg = rows.Scan(&dbName)
		if errMsg != nil {
			return nil, errMsg
		}

		result = append(result, dbName)
	}

	return result, nil
}

// 获取所有表名
// 返回值:
// []string:表名集合
func (this *Connection) GetTableNames(dbName string) ([]string, error) {
	executeSql := "SELECT TABLE_NAME FROM `tables` WHERE table_schema=?"
	rows, errMsg := this.connectionObj.Query(executeSql, dbName)
	if errMsg != nil {
		return nil, fmt.Errorf("获取数据库所有表失败:%v", errMsg.Error())
	}

	result := make([]string, 0)
	var tbName string
	for rows.Next() {
		errMsg = rows.Scan(&tbName)
		if errMsg != nil {
			return nil, errMsg
		}

		result = append(result, tbName)
	}

	return result, nil
}

// 加载数据库
// dbName:数据库名
// 返回值:
// *DbInfo:数据库对象
// error:错误信息
func (this *Connection) LoadDb(dbName string) (*DbInfo, error) {
	return newDbInfo(this, dbName)
}

// 加载指定数据库表
// dbName:数据库名
// tableName:表名
// 返回值:
// *DbTable:数据库对象
// error:错误信息
func (this *Connection) LoadTable(dbName string, tableName string) (*DbTable, error) {
	executeSql := "SELECT TABLE_NAME,TABLE_COMMENT FROM tables WHERE table_schema=? and TABLE_NAME=?"
	row := this.connectionObj.QueryRow(executeSql, dbName, tableName)

	var (
		comment string
		columns []*DbColumn
	)

	errMsg := row.Scan(&tableName, &comment)
	if errMsg != nil {
		return nil, errMsg
	}

	columns, errMsg = getALLColumns(this.connectionObj, dbName, tableName)
	if errMsg != nil {
		return nil, errMsg
	}

	return newDbTable(tableName, comment, columns), nil
}

// 加载指定数据库表
// dbName:数据库名
// tableNames:表名
// 返回值:
// []*DbTable:数据库对象
// error:错误信息
func (this *Connection) LoadTables(dbName string, tableNames ...string) ([]*DbTable, error) {
	result := make([]*DbTable, 0)
	for _, tableName := range tableNames {
		tableItem, errMsg := this.LoadTable(dbName, tableName)
		if errMsg != nil {
			return nil, errMsg
		}

		result = append(result, tableItem)
	}

	return result, nil
}

// 新建一个连接对象
// _userName:用户名
// _passward:密码
// _ip:IP
// _port:端口
// 返回值:
// *DbInfo:数据库对象
// error:错误信息
func NewConnection(_userName string, _passward string, _ip string, _port int) (*Connection, error) {
	// 连接到数据库
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/information_schema?charset=utf8",
		_userName, _passward, _ip, _port)
	_db, errMsg := sql.Open("mysql", connectionStr)
	if errMsg != nil {
		return nil, errMsg
	}

	if errMsg = _db.Ping(); errMsg != nil {
		return nil, errMsg
	}

	return &Connection{
		ip:            _ip,
		port:          _port,
		userName:      _userName,
		passward:      _passward,
		connectionObj: _db,
	}, nil
}
