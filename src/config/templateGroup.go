package config

// 模板配置项
type TemplateItem struct {
	TemplatePath   string `xml:"TemplatePath,attr"`
	FileNameFormat string `xml:"FileNameFormat,attr"`
}

// 分组信息项
type TemplateGroupItem struct {
	// 分组名
	Name string `xml:"Name,attr"`

	// 模板列表
	TemplateList []*TemplateItem `xml:"TemplateItem"`
}
