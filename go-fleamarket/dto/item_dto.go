package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=99999"`
	Description string `json:"description"`
}

// pointer 型にすることで　nil を許容するようにする
// java で null を許容するために参照型にすることと同じ
type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=99999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
