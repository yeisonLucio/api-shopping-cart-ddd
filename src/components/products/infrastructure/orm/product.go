package product_orm

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	Description string `gorm:"size:255"`
	Price       string `gorm:"size:50"`
	Quantity    string `gorm:"size:50"`
}

func (p *Product) Parse(structure interface{}) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &structure)
}

func (Product) TableName() string {
	return "products"
}
