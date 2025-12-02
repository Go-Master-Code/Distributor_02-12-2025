package model

import "time"

type Ongkir struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Ongkir) TableName() string {
	return "ongkir"
}
