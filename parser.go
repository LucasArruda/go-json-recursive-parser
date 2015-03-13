package jsonparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var shouldAppendJSON bool = false

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func RequestUrlWithoutJSON(uri string) {
	shouldAppendJSON = true
	Request(uri)
}

func Request(uri string) {
	if (shouldAppendJSON) {
		uri = appendJSON(uri)
	}
	data := getUrlData(uri)

	if val, ok := data["follow"]; ok {
		Request(val.(string))
	} else {
		for _, val := range data {
			fmt.Println(val)
		}
	}
}

func appendJSON(uri string) string {
	u, err := url.Parse(uri)
    check(err)
    u.Path += ".json"

	return u.String()
}

func getUrlData(uri string) map[string]interface{} {
	return getData(getUrl(uri))
}

func getUrl(uri string) []byte {
	res, err := http.Get(uri)
	check(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	check(err)

	return body
}

func getData(body []byte) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data
}
