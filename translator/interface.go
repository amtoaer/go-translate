package translator

import (
	"github.com/amtoaer/go-translate/config"
	"github.com/amtoaer/go-translate/utils"
)

type Method uint8

const (
	BAIDU Method = iota
	NIU
	YOUDAO
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
	switch m {
	case BAIDU:
		return &baidu{token: utils.CheckEmpty(config.GlobalConf.BaiduToken, config.BAIDU_DEFAULT_TOKEN)}
	case NIU:
		return &niu{
			Mapper: Mapper{map[string]string{
				"zh": "zh",
				"en": "en",
			}},
			apiKey: utils.CheckEmpty(config.GlobalConf.NiuToken, config.NIU_DEFAULT_TOKEN),
		}
	default:
		return &youdao{
			id:     utils.CheckEmpty(config.GlobalConf.YoudaoID, config.YOUDAO_DEFAULT_ID),
			secret: utils.CheckEmpty(config.GlobalConf.YoudaoSecret, config.YOUDAO_DEFAULT_SECRET),
			Mapper: Mapper{map[string]string{
				"zh": "zh-CHS",
				"en": "en",
			}},
		}
	}
}
