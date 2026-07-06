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

type Symbole struct {
	Name              string  `gorm:"column:name"`
	SharesOutstanding float64 `gorm:"column:shares_outstanding"`
}

func (TickertHistory) TableName() string {
	return "tickert_history"
}

func (Symbole) TableName() string {
	return "symbols"
}

type PriceLevel struct {
	ClosePrice float64 `gorm:"colume:close_price"`
	Volume     float64 `gorm:"colume:volume"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&TickertHistory{}, &Symbole{})

	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var result []string

		err := db.Model(&TickertHistory{}).Distinct("name").Pluck("name", &result).Error

		if err != nil {
			log.Fatalln(err)
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

	http.HandleFunc("/constituents", func(w http.ResponseWriter, r *http.Request) {
		var constituents []Symbole
		db.Model(&Symbole{}).Select("name", "shares_outstanding").Scan(&constituents)
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().
			Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(constituents); err != nil {
			log.Println("Error in json Encoding", err)
		}
	})

	http.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		var priceLevel []PriceLevel
		err := db.Model(&TickertHistory{}).
			Select("close_price", "volume").
			Where("name=?", name).
			Order("date_price Asc").
			Scan(&priceLevel)
		if err.Error != nil {
			log.Println("Error", err.Error)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return

		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(priceLevel); err != nil {
			log.Println("Error in json Encoding", err)
		}
	})

	http.HandleFunc("/out_standing_shares", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		}
		var result []int
		err := db.Model(&TickertHistory{}).Select()
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error", err)
	}
}
