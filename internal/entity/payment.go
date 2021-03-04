package entity

import "time"

// Payment type
type Payment struct {
	ID       uint32
	UserID   uint32
	Category string
	Price    uint32
	Date     time.Time
	Memo     string
}
