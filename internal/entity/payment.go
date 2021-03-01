package entity

import "time"

// Payment カテゴリのモデル
type Payment struct {
	ID         uint32
	UserID     uint32
	CategoryID uint32
	Price      uint32
	Date       time.Time
	memo       *string
	Bys
	Ats
}
