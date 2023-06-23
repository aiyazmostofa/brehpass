package login

import (
	"fmt"
	"github.com/aiyazmostofa/brehpass/internal/base"
	"net/http"
)

var API = base.API{
	Path:   "/login",
	Create: base.MethodNotAllowedHandler,
	Read:   read,
	Update: base.MethodNotAllowedHandler,
	Delete: base.MethodNotAllowedHandler,
}

func read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<h1>Hello</h1>")
}
