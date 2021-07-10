package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
)

var GlobalConf *Config = &Config{}

const (
	LANG_ERROR    = "指定的语言错误，仅支持zh/en"
	REQUEST_ERROR = "网络请求错误，请检查网络状态"
)

const (
	BAIDU_DEFAULT_TOKEN = ""
	NIU_DEFAULT_TOKEN   = "375411e1d55cacc2b850011bb74ae986"
)

type Config struct {
	BaiduToken string
	NiuToken   string
}

func InitConfig() error {

	homeDir, err := homedir.Expand("~/.gotrans")
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(homeDir)
	if err != nil {
		return nil
	}
	return json.Unmarshal(content, GlobalConf)
}
