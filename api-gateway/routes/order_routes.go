package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func RegisterOrderRoutes(router *http.ServeMux) {
	orderServiceURL, _ := url.Parse(os.Getenv("ORDER_SERVICE_URL"))

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(orderServiceURL)
		proxy.ServeHTTP(w, r)
	})

	router.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(orderServiceURL)
		proxy.ServeHTTP(w, r)
	})
}
