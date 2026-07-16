package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TranscationtHistory struct {
	Name       string  `gorm:"column:Name"`
	OpenPrice  float64 `gorm:"column:OpenPrice"`
	LowPrice   float64 `gorm:"column:LowPrice"`
	HighPrice  float64 `gorm:"column:HighPrice"`
	ClosePrice float64 `gorm:"column:ClosePrice"`
	Volume     float64 `gorm:"column:Volume"`
	DatePrice  string  `gorm:column:DatePrice`
}

func (TranscationtHistory) TableName() string {
	return "TranscationtHistory"
}

type Symbols struct {
	Name              string  `gorm:"column:Name"`
	SharesOutstanding float64 `gorm:"column:SharesOutstanding"`
}

func (Symbols) TableName() string {
	return "Symbols"
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&TranscationtHistory{}, &Symbols{})
	if err != nil {
		log.Fatalln(err)
	}
}
