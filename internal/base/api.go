package base

import (
	"net/http"
)

type API struct {
	Path   string
	Create func(w http.ResponseWriter, r *http.Request)
	Read   func(w http.ResponseWriter, r *http.Request)
	Update func(w http.ResponseWriter, r *http.Request)
	Delete func(w http.ResponseWriter, r *http.Request)
}

func InitializeAPI(routes []API) {
	handlerInitializer := func(route API) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				route.Create(w, r)
			} else if r.Method == "GET" {
				route.Read(w, r)
			} else if r.Method == "PUT" {
				route.Update(w, r)
			} else if r.Method == "DELETE" {
				route.Delete(w, r)
			} else {
				MethodNotAllowedHandler(w, r)
			}
		})

		http.Handle(route.Path, handler)
	}

	for _, route := range routes {
		handlerInitializer(route)
	}
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
