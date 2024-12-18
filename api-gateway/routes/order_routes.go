package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func RegisterOrderRoutes(router *http.ServeMux) {
	err := godotenv.Load()
	if (err != nil) {
		log.Fatal("Error loading .env file")
	}

	orderServiceURL, err := url.Parse(os.Getenv("ORDER_SERVICE_URL"))
	if err != nil || orderServiceURL.Scheme == "" || orderServiceURL.Host == "" {
		log.Fatalf("Invalid ORDER_SERVICE_URL: %v", err)
	}

	router.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(orderServiceURL)
		r.URL.Path = "/orders"
		proxy.ServeHTTP(w, r)
	})
}
