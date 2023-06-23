package main

import (
	"net/http"

	"github.com/aiyazmostofa/brehpass/internal/login"
	"github.com/aiyazmostofa/brehpass/internal/base"
)

func main() {
	base.InitializeAPI([]base.API{login.API})
	http.ListenAndServe(":8080", nil)
}
