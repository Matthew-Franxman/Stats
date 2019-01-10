package main

import "fmt"

//an interface tells all of the structs in a package that they must have specific methods associated with them
//the specific interface 'Player' requires all structs to have a 'returnStats()' method

//Player interface to kee track of player cards
type Player interface {
	returnStats() string
}

//a struct is virtually just a collection of varibales with associated methods
//for example Baseball has four variables associated with it

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

//each of these structs have their own 'retrunStats()' method associated with them
//since each of them has different
func (b Baseball) returnStats() {
	fmt.Printf("Name: %s\nNumber: %s\nBatting Average: %f\nRBI: %d", b.Name, b.Number, b.BattingAve, b.RBI)
}

func (b Basketball) returnStats() {
	fmt.Printf("Name: %s\nNumber: %s\nAverage Number of Points: %f\nAverage Number of Rebounds: %f", b.Name, b.Number, b.AvePoints, b.AveRebounds)
}

func (f Football) returnStats() {
	fmt.Printf("Name: %s\nNumber: %s\nTouchdowns: %d\nYards: %d", f.Name, f.Number, f.Touchdowns, f.Yards)
}
