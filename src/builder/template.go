package builder

import (
	"GoCodeGenerator/src/templateUtil"
	"bytes"
	"io/ioutil"
	"text/template"
)

// 获取模版处理结果
// templateName:模版名
// context:上下文信息
// 返回值:
// string:结果字符串
func GetReuslt(templteContent string, context interface{}) (string, error) {
	templateObj, errMsg := getTemplate(templteContent)
	if errMsg != nil {
		return "", errMsg
	}

	buffer := bytes.NewBuffer(nil)

	// 得到模板转换结果
	errMsg = templateObj.Execute(buffer, context)
	if errMsg != nil {
		return "", errMsg
	}

	return buffer.String(), nil
}

// 获取模版处理结果
// templateName:模版名
// context:上下文信息
// 返回值:
// string:结果字符串
func GetReusltByPath(templtePath string, context interface{}) (string, error) {
	templateContent, errMsg := ioutil.ReadFile(templtePath)
	if errMsg != nil {
		return "", errMsg
	}

	return GetReuslt(string(templateContent), context)
}

// 获取模板对象
// templteContent:模板内容
// *template.Template:模板对象
// error:错误信息
func getTemplate(templteContent string) (*template.Template, error) {
	templateObj := template.New("template")

	// 注册处理函数
	funcs := template.FuncMap(templateUtil.GetTemplateFunData())
	templateObj = templateObj.Funcs(funcs)

	// 模板转换
	var errMsg error
	templateObj, errMsg = templateObj.Parse(templteContent)
	if errMsg != nil {
		return nil, errMsg
	}

	return templateObj, errMsg
}

// 获取模板对象
// templteContent:模板内容
// *template.Template:模板对象
// error:错误信息
func getTemplateByPath(templatePath string) (*template.Template, error) {
	templateObj := template.New(templatePath)
	// 注册处理函数
	funcs := template.FuncMap(templateUtil.GetTemplateFunData())
	templateObj = templateObj.Funcs(funcs)

	// 模板转换
	templateObj, errMsg := templateObj.ParseFiles(templatePath)
	if errMsg != nil {
		return nil, errMsg
	}

	return templateObj, errMsg
}
