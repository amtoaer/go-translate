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
	YOUDAO_ERROR  = "翻译失败，请前往有道文本翻译文档根据错误码%s查看具体错误信息"
)

const (
	BAIDU_DEFAULT_TOKEN   = ""
	NIU_DEFAULT_TOKEN     = "375411e1d55cacc2b850011bb74ae986"
	YOUDAO_DEFAULT_ID     = "7320e39925cb0de4"
	YOUDAO_DEFAULT_SECRET = "cGqUsf555c1BNIdhrB6ekB4ouF8MZ380"
)

type Config struct {
	BaiduToken   string
	NiuToken     string
	YoudaoID     string
	YoudaoSecret string
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
