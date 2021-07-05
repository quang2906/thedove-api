package model

type Cart struct {
	Id         int64   `json:"id" gorm:"primaryKey"`
	TotalPrice float64 `json:"totalPrice"`
	Quantity   int64   `json:"quantity"`
	CreatedAt  int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
