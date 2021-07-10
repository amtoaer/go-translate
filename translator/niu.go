package translator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/amtoaer/go-translate/config"
)

type niu struct {
	apiKey string
	Mapper
}

func (n *niu) Translate(content, fromLang, toLang string) string {
	fromLang, toLang, ok := n.checkLang(fromLang, toLang)
	if !ok {
		return config.LANG_ERROR
	}
	form := url.Values{}
	form.Add("from", fromLang)
	form.Add("to", toLang)
	form.Add("apikey", n.apiKey)
	form.Add("src_text", content)
	resp, err := http.PostForm("https://api.niutrans.com/NiuTransServer/translation", form)
	if err != nil {
		log.Fatal(config.REQUEST_ERROR)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var jsonResult map[string]string
	json.Unmarshal(result, &jsonResult)
	if errMsg, ok := jsonResult["error_msg"]; ok {
		return errMsg
	}
	return jsonResult["tgt_text"]
}
