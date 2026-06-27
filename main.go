package main

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TickertHistory struct {
	Name       string    `gorm:"name"`
	OpenPrice  float64   `gorm:"open_price"`
	LowPrice   float64   `gorm:"low_price"`
	HighPrice  float64   `gorm:"high_price"`
	ClosePrice float64   `gorm:"close_price"`
	Volume     float64   `gorm:"volume"`
	DatePrice  time.Time `gorm:"date_price"`
}

func (TickertHistory) TableName() string {
	return "tickert_history"
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&TickertHistory{})
	var result []TickertHistory
	err = db.Select("name").Find(&result).Error
	if err != nil {
		log.Fatalln(err)
	}

}
