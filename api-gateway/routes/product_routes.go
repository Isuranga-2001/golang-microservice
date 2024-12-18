package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func RegisterProductRoutes(router *http.ServeMux) {
	err := godotenv.Load()
	if (err != nil) {
		log.Fatal("Error loading .env file")
	}

	productServiceURL, err := url.Parse(os.Getenv("PRODUCT_SERVICE_URL"))
	if err != nil || productServiceURL.Scheme == "" || productServiceURL.Host == "" {
		log.Fatalf("Invalid PRODUCT_SERVICE_URL: %v", err)
	}

	router.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(productServiceURL)
		r.URL.Path = "/products"
		proxy.ServeHTTP(w, r)
	})
}
