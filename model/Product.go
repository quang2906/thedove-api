package model

type Product struct {
	Id         int64   `json:"id" gorm:"primaryKey"`
	CategoryId int64   `json:"categoryId"`
	Image      []Image `json:"image" gorm:"foreignKey:ProductId"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	IsSale     bool    `json:"isSales"`
	CreatedAt  int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
