package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TickertHistory struct {
	Name       string    `gorm:"column:name;index"`
	OpenPrice  float64   `gorm:"column:open_price"`
	LowPrice   float64   `gorm:"column:low_price"`
	HighPrice  float64   `gorm:"column:high_price"`
	ClosePrice float64   `gorm:"column:close_price"`
	Volume     float64   `gorm:"column:volume"`
	DatePrice  time.Time `gorm:"column:date_price"`
}

func (TickertHistory) TableName() string {
	return "tickert_history"
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&TickertHistory{})

	if err != nil {
		log.Fatalln(err)
	}

	var result []string

	err = db.Model(&TickertHistory{}).Distinct("name").Pluck("name", &result).Error

	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		respone := struct {
			Count  int
			Ticker []string
		}{
			Count:  len(result),
			Ticker: result,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(respone); err != nil {
			log.Println("Error in json Encoding:", err)
		}

	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Failed to start server:", err)
	}

}
