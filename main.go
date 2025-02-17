package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"receipt-processor/controllers"
	"receipt-processor/storage"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", controllers.ProcessReceiptHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", controllers.GetPointsHandler).Methods("GET")

	router.Use(loggingMiddleware)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		storage.PrintStore()
	})
}
