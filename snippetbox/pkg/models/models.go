package models
 
import (
	"errors"
	"time"
)
 
var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Perosn struct {
	ID int
	Name string
	Age int 
	Gender string
	Address string
}

type Menu struct {
	ID int
	Name string
	Price uint
}

type Order struct {
	ID      int
	Surname   string
	Name string
	Amount uint
	Order_date time.Time
}