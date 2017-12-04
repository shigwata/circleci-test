package hello

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Hello World!")
}
