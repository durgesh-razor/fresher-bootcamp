package Models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderID   uint64 `json:"orderID"`
	User      User   `json:"user" binding:"required" gorm:"foreignkey:UserID"`
	UserID    uint64
	Product   Product `json:"product" binding:"required" gorm:"foreignkey:ProductID"`
	ProductID uint64
	Status    string `json:"status"`
}
