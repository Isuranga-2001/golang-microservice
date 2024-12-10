package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func RegisterProductRoutes(router *http.ServeMux) {
	productServiceURL, _ := url.Parse("http://localhost:8001")

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(productServiceURL)
		proxy.ServeHTTP(w, r)
	})

	router.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(productServiceURL)
		proxy.ServeHTTP(w, r)
	})
}
