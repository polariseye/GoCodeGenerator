package config

// 连接配置项
type ConnectionItem struct {
	// 连接名
	Name string `xml:",attr"`

	// 数据库的用户名
	UserName string `xml:",attr"`

	// 密码
	Pwd string `xml:",attr"`

	// IP
	Ip string `xml:",attr"`

	// 端口
	Port int `xml:",attr"`

	// 默认数据库
	DefaultDb string `xml:",attr"`
}
