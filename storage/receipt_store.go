package storage

import (
	"log"
	"sync"
)

var (
	receiptStore = make(map[string]int)
	storeMutex   = sync.Mutex{}
)

func SaveReceipt(id string, points int) {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	receiptStore[id] = points
}

func GetReceiptPoints(id string) (int, bool) {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	points, ok := receiptStore[id]
	return points, ok
}

func PrintStore() {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	if len(receiptStore) == 0 {
		log.Println("Receipt Store is empty")
		return
	}

	log.Println("Current Receipt Store:")
	for id, points := range receiptStore {
		log.Printf("  - ID: %s, Points: %d", id, points)
	}
}
