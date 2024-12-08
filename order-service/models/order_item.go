package models

type OrderItem struct {
	ID        int `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	Quantity  int `json:"quantity"`
	ProductID int `json:"product_id"`
	Order     Order
}