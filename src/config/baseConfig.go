package config

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
)

// 基本配置
type BaseConfig struct {
	// 默认配置文件加载
	DefaultConfig string `xml:DefaultConfig`

	// 头部配置信息
	HeaderInfo *HeaderInfo `xml:HeaderInfo`

	// 数据库连接配置项
	DbConnection []*ConnectionItem `xml:"DbConnection>ConnectionItem"`
}

var (
	// 基础配置
	baseConfigObj *BaseConfig
)

// 初始化
func init() {
	err := loadBaseConfig("config/baseconfig.xml")
	if err != nil {
		panic(err)
	}
}

// 获取配置信息项
// 返回值:
// *BaseConfig:配置数据
func GetBaseConfig() *BaseConfig {
	if baseConfigObj == nil {
		panic(errors.New("基本配置未加载"))
	}
	return baseConfigObj
}

// 从指定配置文件加载配置
// filePath:文件路径
// 返回值:
// error:错误信息
func loadBaseConfig(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	configObj := BaseConfig{}
	err = xml.Unmarshal(data, &configObj)
	if err != nil {
		return err
	}

	baseConfigObj = &configObj

	return nil
}
