package api

import (
	"fmt"
	"github.com/gsxhnd/owl"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	conf, _ := owl.GetByKey("/conf/gateway.yaml")
	fmt.Println(conf)

	_, _ = fmt.Fprint(w, conf)
}
