package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Car struct {
	CarID string          `json:"car_id" gorm:"primaryKey"`
	Name  string          `json:"name_id"`
	Price decimal.Decimal `json:"price"`
}

func (r *Car) BeforeCreate(tx *gorm.DB) (err error) {
	r.CarID = uuid.New().String()
	return
}
