package config

// 文件头信息
type HeaderInfo struct {
	// 作者名
	AuthorName string `xml:"AuthorName,attr"`

	// 时间格式
	TimeFormat string `xml:"TimeFormat,attr"`
}
