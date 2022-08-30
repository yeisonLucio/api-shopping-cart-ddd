package domain

type Filters struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       string `json:"price" form:"price"`
}
