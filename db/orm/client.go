package orm

import (
	"louis/pw/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbClient struct {
	conn *gorm.DB
}

func (db *dbClient) GetAllClient() ([]*model.Client, error) {
	var cl []*model.Client
	err := db.conn.Find(&cl).Error
	if err != nil {
		return nil, err
	}
	return cl, nil
}

func (db *dbClient) CreateClient(c *model.Client) (*model.Client, error) {
	c.ID = uuid.New().String()
	c.Valid = false
	c.IdCardPath = ""
	c.IdCardDate = ""
	c.IdCardNum = ""
	c.IdCardType = 0
	c.IdCardAuthority = ""

	err := db.conn.Create(c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (db *dbClient) DeleteClient(id string) error {
	var c model.Client
	return db.conn.Where("id = ?", id).Delete(&c).Error
}

func (db *dbClient) UpdateClient(id string, data map[string]interface{}) (*model.Client, error) {
	delete(data, "valid")
	err := db.conn.Model(&model.Client{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetClientById(id)
}

func (db *dbClient) ValidateClient(id string) (*model.Client, error) {
	err := db.conn.Model(&model.Client{}).Where("id = ?", id).Update("Valid", true).Error
	if err != nil {
		return nil, err
	}
	return db.GetClientById(id)
}

func (db *dbClient) UnvalidateClient(id string) (*model.Client, error) {
	err := db.conn.Model(&model.Client{}).Where("id = ?", id).Update("Valid", false).Error
	if err != nil {
		return nil, err
	}
	return db.GetClientById(id)
}

func (db *dbClient) GetClientById(id string) (*model.Client, error) {
	var c model.Client
	err := db.conn.First(&c, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}
