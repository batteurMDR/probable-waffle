package orm

import (
	"louis/pw/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbProduct struct {
	conn *gorm.DB
}

func (db *dbProduct) GetAllProduct() ([]*model.Product, error) {
	var pr []*model.Product
	err := db.conn.Find(&pr).Error
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func (db *dbProduct) CreateProduct(p *model.Product) (*model.Product, error) {
	p.ID = uuid.New().String()
	err := db.conn.Create(p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (db *dbProduct) DeleteProduct(id string) error {
	var p model.Product
	return db.conn.Where("id = ?", id).Delete(&p).Error
}

func (db *dbProduct) UpdateProduct(id string, data map[string]interface{}) (*model.Product, error) {
	err := db.conn.Model(&model.Product{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetProductById(id)
}

func (db *dbProduct) GetProductById(id string) (*model.Product, error) {
	var p model.Product
	err := db.conn.First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
