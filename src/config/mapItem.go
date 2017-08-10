package config

// 信息映射项
type MapItem struct {
	// 老的值
	OldValue string `xml:",attr"`

	// 目标值
	TargetValue string `xml:",attr"`

	// 导入的命名空间
	Import string `xml:",attr"`
}
