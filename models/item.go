package models

type Item struct {
    ID    uint    `json:"id" gorm:"primaryKey"`
    Name  string  `json:"name"`
    Stock int     `json:"stock"`
    Price float64 `json:"price"`
}
