package translator

type Method uint8

const (
	BAIDU Method = iota
	GOOGLE
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
	switch m {
	case BAIDU:
		return &baidu{}
	case GOOGLE:
		return &google{}
	case NIU:
		return &niu{Mapper: Mapper{map[string]string{
			"zh": "zh",
			"en": "en",
		}}}
	default:
		return
	}
}
