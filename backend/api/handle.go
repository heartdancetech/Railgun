package api

import (
	"encoding/json"
	"github.com/gsxhnd/owl"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetKeysHandle(w http.ResponseWriter, req *http.Request) {
	keys, _ := owl.GetRemoteKeys("/conf/")
	SendRes(w, nil, keys)
}

func GetValueHandle(w http.ResponseWriter, req *http.Request) {
	params, _ := url.ParseQuery(req.URL.RawQuery)
	key := params.Get("key")
	v, err := owl.GetRemote(key)
	if err != nil {
		SendRes(w, err, nil)
	} else {
		SendRes(w, nil, v)
	}
}

func SaveValueHandle(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var params struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	_ = json.Unmarshal(body, &params)
	err := owl.PutRemote(params.Key, params.Value)
	if err != nil {
		SendRes(w, err, nil)
	} else {
		SendRes(w, nil, nil)
	}
}
