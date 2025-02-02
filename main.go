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

//curl -X POST http://localhost:8080/receipts/process \
//-H "Content-Type: application/json" \
//-d '{
//"retailer": "M&M Corner Market",
//"purchaseDate": "2022-03-20",
//"purchaseTime": "14:33",
//"items": [
//{
//"shortDescription": "Gatorade",
//"price": "2.25"
//},{
//"shortDescription": "Gatorade",
//"price": "2.25"
//},{
//"shortDescription": "Gatorade",
//"price": "2.25"
//},{
//"shortDescription": "Gatorade",
//"price": "2.25"
//}
//],
//"total": "9.00"
//}'
