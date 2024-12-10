package models

type Order struct {
	ID            int         `json:"order_id" gorm:"primaryKey;autoIncrement"`
	UserID        int         `json:"user_id"`
	PaymentStatus string      `json:"payment_status"`
	PaymentMethod string      `json:"payment_method"`
	OrderItems    []OrderItem `gorm:"foreignKey:OrderID"` // Define foreign key
}