package config

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
)

// 基本配置
type TemplateConfig struct {

	// 数据类型映射表
	DataTypeMapTable []*MapItem `xml:"DataTypeMapTable>MapItem"`

	// 列名称映射表
	ColumnMapTable []*MapItem `xml:"ColumnMapTable>MapItem"`

	// 表前缀处理表
	TablePrefix []*FixionItem `xml:"TablePrefix>FixionItem"`

	// 表名后缀处理表
	TableStuffix []*FixionItem `xml:"TableStuffix>FixionItem"`

	// 模板配置信息
	TemplateGroup []*TemplateGroupItem `xml:"Template>GroupItem"`
}

var (
	templateConfigObj *TemplateConfig
)

// 获取配置信息项
// 返回值:
// *TemplateConfig:配置数据
func GetTemplateConfig() *TemplateConfig {
	if templateConfigObj == nil {
		panic(errors.New("基本配置未加载"))
	}
	return templateConfigObj
}

// 从指定配置文件加载配置
// filePath:文件路径
// 返回值:
// error:错误信息
func LoadTemplateConfig(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	configObj := TemplateConfig{}
	err = xml.Unmarshal(data, &configObj)
	if err != nil {
		return err
	}

	templateConfigObj = &configObj

	return nil
}
