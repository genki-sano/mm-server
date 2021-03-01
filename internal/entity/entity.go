package entity

import "time"

// CreatedBy 作成者のusersテーブルのID(PK)
type CreatedBy uint32

// UpdatedBy 更新者のusersテーブルのID(PK)
type UpdatedBy uint32

// CreatedAt 作成日時
type CreatedAt time.Time

// UpdatedAt 更新日時
type UpdatedAt time.Time

// Bys type
type Bys struct {
	CreatedBy CreatedBy
	UpdatedBy UpdatedBy
}

// Ats type
type Ats struct {
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
}
