package freegufo_template

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func init() {
	http.HandleFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "handler")
}
