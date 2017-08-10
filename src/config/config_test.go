package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConfig(context *testing.T) {
	err := loadConfig("config.xml")
	if err != nil {
		context.Error(err)
		return
	}

	val, err := json.Marshal(GetConfig())
	if err != nil {
		context.Error(err)
		return
	}

	fmt.Printf("\r\n当前的数据为:%s\r\n", string(val))
}
