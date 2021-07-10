package translator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/amtoaer/go-translate/config"
	"github.com/amtoaer/go-translate/utils"
)

type youdao struct {
	id, secret string
	Mapper
}

func (y *youdao) Translate(content, fromLang, toLang string) string {
	fromLang, toLang, ok := y.checkLang(fromLang, toLang)
	if !ok {
		return config.LANG_ERROR
	}
	var (
		salt        string = utils.GenerateUUID()
		currentTime string = fmt.Sprintf("%d", time.Now().Unix())
		sign        string = utils.GenerateYoudaoSign(y.id, content, salt, currentTime, y.secret)
	)
	form := url.Values{}
	form.Add("q", content)
	form.Add("from", fromLang)
	form.Add("to", toLang)
	form.Add("appKey", y.id)
	form.Add("salt", salt)
	form.Add("sign", sign)
	form.Add("signType", "v3")
	form.Add("curtime", currentTime)
	resp, err := http.PostForm("https://openapi.youdao.com/api", form)
	if err != nil {
		return config.REQUEST_ERROR
	}
	result, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var jsonResult map[string]interface{}
	json.Unmarshal(result, &jsonResult)
	errCode := jsonResult["errorCode"].(string)
	if errCode != "0" {
		return fmt.Sprintf(config.YOUDAO_ERROR, errCode)
	}
	return jsonResult["translation"].([]interface{})[0].(string)
}
