package infrastructure

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/orm"
	"github.com/yeisonLucio/shopping-cart/src/database/connections"
)

type MysqlProduct struct {
}

func NewMysqlProduct() *MysqlProduct {
	return &MysqlProduct{}
}

func (mp *MysqlProduct) CreateProduct(product domain.Product) (bool, error) {
	var newProduct orm.Product

	if err := product.Parse(&newProduct); err != nil {
		return false, err
	}

	if result := connections.DB.Create(&newProduct); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (mp *MysqlProduct) GetProductById(productID uint) (domain.Product, error) {

	var productDB orm.Product
	var product domain.Product

	if res := connections.DB.First(&productDB, productID); res.Error != nil {
		return product, res.Error
	}

	if err := productDB.Parse(&product); err != nil {
		return product, err
	}

	return product, nil
}

func (mp *MysqlProduct) Save(product domain.Product) error {

	var productToUpdate orm.Product

	if err := product.Parse(&productToUpdate); err != nil {
		return err
	}

	if result := connections.DB.UpdateColumns(&productToUpdate); result.Error != nil {
		return result.Error
	}

	return nil
}

func (mp *MysqlProduct) GetFilteredProducts(filters domain.Filters) ([]domain.Product, error) {

	var products []domain.Product
	result := connections.DB.
		Where("name LIKE ?", "%"+filters.Name+"%").
		Where("description LIKE ?", "%"+filters.Description+"%").
		Where("Price LIKE ?", "%"+filters.Price+"%").
		Where("deleted_at is NULL").
		Find(&products)

	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (mp *MysqlProduct) DeleteProduct(productID uint) error {
	return connections.DB.
		Where("id = ?", productID).
		Delete(&orm.Product{}).Error
}
