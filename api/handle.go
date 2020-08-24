package api

import (
	"fmt"
	"github.com/gsxhnd/owl"
	"net/http"
)

func GetKeysHandle(w http.ResponseWriter, req *http.Request) {
	keys, _ := owl.GetKeys("/conf")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, keys)
}
