package model

type Price struct {
	Id         int64   `json:"id" gorm:"primaryKey"`
	ProductId  int64   `json:"product_id"`
	Value      float64 `json:"value"`
	CreatedAt  int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
