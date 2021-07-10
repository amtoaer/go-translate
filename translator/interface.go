package translator

import (
	"github.com/amtoaer/go-translate/config"
	"github.com/amtoaer/go-translate/utils"
)

type Method uint8

const (
	BAIDU Method = iota
	NIU
)

type Translator interface {
	Translate(content, fromLang, toLang string) string
}

type Mapper struct {
	mapper map[string]string
}

func (m *Mapper) checkLang(fromLang, toLang string) (string, string, bool) {
	validFromLang, isFromOk := m.mapper[fromLang]
	validToLang, isToOk := m.mapper[toLang]
	return validFromLang, validToLang, isFromOk && isToOk
}

func NewTranslator(m Method) (t Translator) {
	var token string
	switch m {
	case BAIDU:
		token = utils.CheckEmpty(config.GlobalConf.BaiduToken, config.BAIDU_DEFAULT_TOKEN)
		return &baidu{token: token}
	case NIU:
		token = utils.CheckEmpty(config.GlobalConf.NiuToken, config.NIU_DEFAULT_TOKEN)
		return &niu{
			Mapper: Mapper{map[string]string{
				"zh": "zh",
				"en": "en",
			}},
			apiKey: token,
		}
	default:
		return
	}
}
