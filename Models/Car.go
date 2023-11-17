package Models

import "time"

type Car struct {
	Id           int64     `json:"id,omitempty" gorm:"column:id;primary_key;autoIncrement"`
	Mark         string    `json:"mark,omitempty" gorm:"column:mark"`
	Model        string    `json:"model,omitempty" gorm:"column:model"`
	Year         uint32    `json:"year,omitempty"  gorm:"column:year"`
	Mileage      int64     `json:"mileage,omitempty"  gorm:"column:mileage"`
	Price        int64     `json:"price,omitempty"  gorm:"column:price"`
	Color        string    `json:"color,omitempty"  gorm:"column:color"`
	Transmission string    `json:"transmission,omitempty"  gorm:"column:transmission"`
	FuelType     string    `json:"fuel_type,omitempty"  gorm:"column:fuel_type"`
	CreatedAt    time.Time `json:"created_at"  gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at"  gorm:"column:updated_at"`
	SellerId     int64     `json:"seller_id,omitempty"  gorm:"column:seller_id"`
}
