package config

// 前缀后缀配置项
type FixionItem struct {
	// 前缀后缀值
	Value string `xml:",attr"`

	// 处理方式
	MethodType string `xml:",attr"`

	// 目标值
	TargetValue string `xml:",attr"`
}
