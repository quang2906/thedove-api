package model

type Review struct {
	Id         int64   `json:"id" gorm:"primaryKey"`
	ProductId  int64   `json:"productId"`
	Rating     float64 `json:"rating"`
	CreatedAt  int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
