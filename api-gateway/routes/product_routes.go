package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func RegisterProductRoutes(router *http.ServeMux) {
	productServiceURL, _ := url.Parse(os.Getenv("PRODUCT_SERVICE_URL"))

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(productServiceURL)
		proxy.ServeHTTP(w, r)
	})

	router.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(productServiceURL)
		proxy.ServeHTTP(w, r)
	})
}
