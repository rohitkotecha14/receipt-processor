package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"
	"receipt-processor/storage"
)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	log.Printf("New receipt received for processing: %s", string(bodyBytes))
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Please verify input. Invalid JSON", http.StatusBadRequest)
		return
	}

	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" || len(receipt.Items) == 0 {
		http.Error(w, "Please verify input. Missing required fields", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()

	points, err := services.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, "Error calculating points (Error: "+err.Error()+")", http.StatusBadRequest)
		return
	}

	storage.SaveReceipt(id, points)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, ok := storage.GetReceiptPoints(id)
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
