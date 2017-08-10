package session

import (
	"GoCodeGenerator/src/builder"
	"GoCodeGenerator/src/config"
	"GoCodeGenerator/src/dbHelper"
	"GoCodeGenerator/src/util/uprint"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

const con_BasePath = "config/"

type BuildSession struct {
	nowConfigFileName   string                    //// 是否已经加载配置文件
	connection          *dbHelper.Connection      //// 连接对象
	allDbNames          []string                  //// 此连接下的所有数据库名
	allTableNames       []string                  //// 所有表名
	nowDb               string                    //// 当前查看的数据库名
	nowTables           []string                  //// 当前操作的所有表名
	targetTemplateGroup *config.TemplateGroupItem //// 目标模板分组
}

// 获取所有配置文件名
// 返回值:
// fileNames:所有配置文件名
// err:错误信息
func (this *BuildSession) GetConfigFileNames() (fileNames []string, err error) {
	fileNames = make([]string, 0)
	fileInfo, err := os.Stat(con_BasePath)
	if err != nil {
		return fileNames, err
	}
	if fileInfo.IsDir() == false {
		return fileNames, fmt.Errorf("is not a directory")
	}

	fileObj, err := os.Open(con_BasePath)
	defer fileObj.Close()
	if err != nil {
		return fileNames, fmt.Errorf("open dir error%v\r\n", err)
	}
	allFileInfo, err := fileObj.Readdir(0)
	if err != nil {
		return fileNames, fmt.Errorf("get all config file error%v\r\n", err)
	}

	// 读取所有文件名
	for _, item := range allFileInfo {
		if item.IsDir() {
			continue
		}

		// 只加载指定前缀的文件
		if strings.HasPrefix(item.Name(), "config_") == false {
			continue
		}

		fileNames = append(fileNames, item.Name())
	}

	return
}

// 加载配置文件
// fileName:配置文件路径名
func (this *BuildSession) LoadConfig(fileName string) error {
	fullFileName := path.Join(con_BasePath, fileName)
	_, err := os.Stat(fullFileName)
	if err != nil {
		uprint.Println("file error:", err.Error())
		return err
	}

	err = config.LoadTemplateConfig(fullFileName)
	if err == nil {
		this.nowConfigFileName = fileName
	}

	return err
}

func (this *BuildSession) Open(_userName string, _passward string, _ip string, _port int, _defaultDb string) error {
	if len(this.nowConfigFileName) <= 0 {
		return fmt.Errorf("not load config file")
	}

	_connection, err := dbHelper.NewConnection(_userName, _passward, _ip, _port)
	if err != nil {
		return err
	}

	_dbNames, err := _connection.GetDataBaseNames()
	if err != nil {
		return err
	}

	println("connect success ")
	this.connection = _connection
	this.allDbNames = _dbNames
	this.allTableNames = make([]string, 0)
	this.targetTemplateGroup = nil

	if len(_defaultDb) > 0 {
		println("start load default db:", _defaultDb)
		if this.SelectDb(_defaultDb) == nil {
			println("default db load finish:", _defaultDb)
		}
	}

	return nil
}

func (this *BuildSession) OpenFromConfig(configName string) error {
	if len(this.nowConfigFileName) <= 0 {
		return fmt.Errorf("not load config file")
	}

	configItem, err := getConnectionConfig(configName)
	if err != nil {
		return err
	}

	return this.Open(configItem.UserName, configItem.Pwd, configItem.Ip, configItem.Port, configItem.DefaultDb)
}

func (this *BuildSession) Close() {
	if this.connection == nil {
		return
	}

	this.connection.CloseConnection()

	this.allDbNames = make([]string, 0)
	this.allTableNames = make([]string, 0)
	this.nowTables = make([]string, 0)
	this.connection = nil
}

func (this *BuildSession) QueryConnectionConfig(configName string) []string {
	result := make([]string, 0)
	if len(this.nowConfigFileName) <= 0 {
		return result
	}

	// 设置默认数据库
	configName = strings.TrimSpace(configName)
	if len(configName) <= 0 {
		for _, item := range config.GetBaseConfig().DbConnection {
			result = append(result, item.Name)
		}
		return result
	}

	configName = strings.ToLower(configName)

	for _, item := range config.GetBaseConfig().DbConnection {
		if strings.Contains(strings.ToLower(item.Name), configName) {
			result = append(result, item.Name)
		}
	}

	return result
}

func (this *BuildSession) SelectDb(dbName string) error {
	// 设置默认数据库
	dbName = strings.TrimSpace(dbName)
	if len(dbName) <= 0 {
		return fmt.Errorf("dbName is empty")
	}

	tableList, err := this.connection.GetTableNames(dbName)
	if err != nil {
		return fmt.Errorf("load table error:%v", err.Error())
	}

	this.nowDb = dbName
	this.allTableNames = tableList

	return nil
}

func (this *BuildSession) QueryDb(dbName string) ([]string, error) {
	if this.connection == nil {
		return nil, errors.New("error:no open connection")
	}

	// 设置默认数据库
	dbName = strings.TrimSpace(dbName)
	if len(dbName) <= 0 {
		return this.allDbNames[:], nil
	}

	dbName = strings.ToLower(dbName)

	result := make([]string, 0)
	for _, item := range this.allDbNames {
		if strings.Contains(strings.ToLower(item), dbName) {
			result = append(result, item)
		}
	}

	return result, nil
}

func (this *BuildSession) QueryTable(tableName string) ([]string, error) {
	if len(this.nowDb) <= 0 {
		return nil, fmt.Errorf("no select db")
	}

	// 设置默认数据库
	tableName = strings.TrimSpace(tableName)
	if len(tableName) <= 0 {
		return this.allTableNames[:], nil
	}

	tableName = strings.ToLower(tableName)

	result := make([]string, 0)
	for _, item := range this.allTableNames {
		tmpItem := strings.ToLower(item)
		if strings.Contains(tmpItem, tableName) {
			result = append(result, item)
		}
	}

	return result, nil
}

func (this *BuildSession) QueryTemplateGroup(groupName string) []string {
	result := make([]string, 0)
	// 设置默认数据库
	groupName = strings.TrimSpace(groupName)
	if len(groupName) <= 0 {
		for _, item := range config.GetTemplateConfig().TemplateGroup {
			result = append(result, item.Name)
		}

		return result
	}

	groupName = strings.ToLower(groupName)
	for _, item := range config.GetTemplateConfig().TemplateGroup {
		if strings.Contains(strings.ToLower(item.Name), groupName) {
			result = append(result, item.Name)
		}
	}

	return result
}

func (this *BuildSession) SelectTemplateGroup(groupName string) error {
	template, exist := this.checkIfTemplateGroupExist(groupName)
	if exist == false {
		return fmt.Errorf("templategroup no exist :%v", groupName)
	}

	this.targetTemplateGroup = template
	return nil
}

func (this *BuildSession) Build(savePath string, tableNames ...string) error {
	// 加载分组配置
	if this.targetTemplateGroup == nil {
		return fmt.Errorf("no select template group")
	}

	if len(this.nowDb) <= 0 {
		return fmt.Errorf("no select database")
	}

	// 加载数据库表数据
	targetTables := make([]*dbHelper.DbTable, 0, len(tableNames))
	for _, item := range tableNames {
		resultTable, exist := this.checkIfTableExist(item)
		if exist == false {
			return fmt.Errorf("table no exist: %v", item)
		}

		tableItem, err := this.connection.LoadTable(this.nowDb, resultTable)
		if err != nil {
			return fmt.Errorf("load build table error:%v", err.Error())
		}

		targetTables = append(targetTables, tableItem)
	}

	// 生成
	return builder.Build(targetTables, this.targetTemplateGroup, nil, savePath)
}

func (this *BuildSession) BuildByLike(savePath string, tbNameLike string) error {
	tableList, err := this.QueryTable(tbNameLike)
	if err != nil {
		return err
	}

	return this.Build(savePath, tableList...)
}

func (this *BuildSession) checkIfTableExist(tableName string) (string, bool) {
	tableName = strings.TrimSpace(tableName)
	tableName = strings.ToLower(tableName)

	for _, item := range this.allTableNames {
		if strings.ToLower(item) == tableName {
			return item, true
		}
	}
	return "", false
}

func (this *BuildSession) checkIfTemplateGroupExist(groupName string) (*config.TemplateGroupItem, bool) {
	groupName = strings.TrimSpace(groupName)
	groupName = strings.ToLower(groupName)
	for _, item := range config.GetTemplateConfig().TemplateGroup {
		if groupName == strings.ToLower(item.Name) {
			return item, true
		}
	}

	return nil, false
}

func (this *BuildSession) StateVal() string {
	result := ""
	if len(this.nowConfigFileName) <= 0 {
		result += "configfile:"
	} else {
		result += "configfile:" + this.nowConfigFileName
	}

	if this.connection == nil {
		result += "\r\nconnection:"
	} else {
		result += "\r\nconnection:" + this.connection.String()
	}

	if len(this.nowDb) <= 0 {
		result += "\r\ndb:"
	} else {
		result += "\r\ndb:" + this.nowDb
	}

	if this.targetTemplateGroup == nil {
		result += "\r\ntemplate:"
	} else {
		result += "\r\ntemplate:" + this.targetTemplateGroup.Name
	}

	return result
}

// 重新加载相关信息
func (this *BuildSession) Refresh() error {
	dbnames, err := this.connection.GetDataBaseNames()
	if err != nil {
		return err
	}

	this.allDbNames = dbnames

	ifExist := false
	for _, item := range this.allDbNames {
		if item == this.nowDb {
			this.SelectDb(this.nowDb)
			ifExist = true
			break
		}
	}

	if !ifExist {
		this.nowDb = ""
		this.nowTables = make([]string, 0)
	}

	return nil
}

func NewSession() *BuildSession {
	session := &BuildSession{
		connection:    nil,
		allDbNames:    make([]string, 0),
		allTableNames: make([]string, 0),
		nowTables:     make([]string, 0),
	}

	// 加载默认配置文件
	err := session.LoadConfig(config.GetBaseConfig().DefaultConfig)
	if err == nil {
		fmt.Println("成功加载默认配置文件:", config.GetBaseConfig().DefaultConfig)
	}

	return session
}

func getConnectionConfig(configName string) (*config.ConnectionItem, error) {
	for _, item := range config.GetBaseConfig().DbConnection {
		if strings.ToLower(item.Name) == strings.ToLower(configName) {
			return item, nil
		}
	}

	return nil, fmt.Errorf("connection config no exist")
}
