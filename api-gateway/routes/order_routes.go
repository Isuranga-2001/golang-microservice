package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func RegisterOrderRoutes(router *http.ServeMux) {
	orderServiceURL, _ := url.Parse("http://localhost:8002")

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(orderServiceURL)
		proxy.ServeHTTP(w, r)
	})

	router.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(orderServiceURL)
		proxy.ServeHTTP(w, r)
	})
}
