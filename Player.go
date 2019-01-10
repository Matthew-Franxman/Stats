package main

import "fmt"

//Player interface to kee track of player cards
type Player interface {
	returnStats() string
}

//Baseball player card
type Baseball struct {
	Name       string
	Number     string
	BattingAve float64
	RBI        int
}

//Basketball player card
type Basketball struct {
	Name        string
	Number      string
	AvePoints   float64
	AveRebounds float64
}

//Football Player card
type Football struct {
	Name       string
	Number     string
	Touchdowns int
	Yards      int
}

func (b Baseball) returnStats() {
	fmt.Printf("Name: %s\nNumber: %s\nBatting Average: %f\nRBI: %d", b.Name, b.Number, b.BattingAve, b.RBI)
}

func (b Basketball) returnStats() {
	fmt.Printf("Name: %s\nNumber: %s\nAverage Number of Points: %f\nAverage Number of Rebounds: %f", b.Name, b.Number, b.AvePoints, b.AveRebounds)
}

func (f Football) returnStats() {
	fmt.Printf("Name: %s\nNumber: %s\nTouchdowns: %d\nYards: %d", f.Name, f.Number, f.Touchdowns, f.Yards)
}
