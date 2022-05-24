package orm

import (
	"louis/pw/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbSell struct {
	conn *gorm.DB
}

func (db *dbSell) GetAllSell() ([]*model.Sell, error) {
	var se []*model.Sell
	err := db.conn.Find(&se).Error
	if err != nil {
		return nil, err
	}
	return se, nil
}

func (db *dbSell) CreateSell(s *model.Sell) (*model.Sell, error) {
	s.ID = uuid.New().String()
	err := db.conn.Create(s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (db *dbSell) GetSellById(id string) (*model.Sell, error) {
	var s model.Sell
	err := db.conn.First(&s, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}
