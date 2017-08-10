package builder

import (
	"GoCodeGenerator/src/config"
	"GoCodeGenerator/src/dbHelper"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 使用指定模板组生成模板内容
// tableList:待处理的数据表集合
// templateGroup:模板配置组对象
// extraData:额外的附加模板上下文
// savePath:存储路径
// 返回值:
// error:错误信息
// bool:是否成功
func Build(tableList []*dbHelper.DbTable, templateGroup *config.TemplateGroupItem, extraData map[string]interface{}, savePath string) error {
	// 获取存储路径
	if len(savePath) <= 0 {
		savePath = "resultcode/"
	}

	for _, tableItem := range tableList {
		for _, templateItem := range templateGroup.TemplateList {
			// 组装模板上下文
			context := make(map[string]interface{})
			for key, val := range extraData {
				context[key] = val
			}
			context["TargetTable"] = tableItem
			context["BaseConfig"] = config.GetBaseConfig()
			context["TemplateConfig"] = config.GetTemplateConfig()

			result, errMsg := GetReusltByPath(templateItem.TemplatePath, context)
			if errMsg != nil {
				return errMsg
			}

			targetFileName := tableItem.TableName
			flNameTemplate := strings.TrimSpace(templateItem.FileNameFormat)
			if len(flNameTemplate) > 0 {
				targetFileName, errMsg = GetReuslt(flNameTemplate, context)
				if errMsg != nil {
					return errMsg
				}
			}

			// 获取存储路径
			itemPath := path.Join(savePath, targetFileName)

			// 存储到文件
			saveResult(itemPath, result)
		}
	}

	return nil
}

// 存储结果
// filePath:存储路径
// buffer:字节内容
// 返回值:
// error:错误信息
func saveResult(filePath string, content string) error {
	dir, _ := path.Split(filePath)
	errMsg := os.MkdirAll(dir, os.ModeDir)
	if errMsg != nil {
		return fmt.Errorf("文件夹创建失败:%v", errMsg.Error())
	}

	return ioutil.WriteFile(filePath, []byte(content), os.ModeTemporary)
}
