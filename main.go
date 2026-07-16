package main

type TranscationtHistory struct {
	Name       string  `grom:"column:Name"`
	OpenPrice  float64 `grom:"column:OpenPrice"`
	LowPrice   float64 `grom:"column:LowPrice"`
	HighPrice  float64 `grom:"column:HighPrice"`
	ClosePrice float64 `grom:"column:ClosePrice"`
	Volume     float64 `grom:"column:Volume"`
	DatePrice  string  `grom:column:DatePrice`
}

func (TranscationtHistory) TableName() string {
	return "TranscationtHistory"
}

type Symbols struct {
	Name              string  `grom:"column:Name"`
	SharesOutstanding float64 `grom:"column:SharesOutstanding"`
}

func (Symbols) TableName() string {
	return "Symbols"
}

func main() {

}
