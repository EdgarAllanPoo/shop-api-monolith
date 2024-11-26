package models

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type Product struct {
	ID          uint    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string  `gorm:"not null" json:"name"`
	Category    string  `gorm:"not null" json:"category"`
	Description string  `gorm:"size:255" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
}
