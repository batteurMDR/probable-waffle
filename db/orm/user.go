package orm

import (
	"louis/pw/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbUser struct {
	conn *gorm.DB
}

func (db *dbUser) GetAllUser() ([]*model.User, error) {
	var us []*model.User
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (db *dbUser) CreateUser(u *model.User) (*model.User, error) {
	err := db.conn.Transaction(func(tx *gorm.DB) error {
		u.ID = uuid.New().String()
		err := db.conn.Create(u).Error
		if err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *dbUser) DeleteUser(id string) error {
	var u model.User
	return db.conn.Where("id = ?", id).Delete(&u).Error
}

func (db *dbUser) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	err := db.conn.Model(&model.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetUserById(id)
}

func (db *dbUser) GetUserById(id string) (*model.User, error) {
	var u model.User
	err := db.conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *dbUser) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := db.conn.First(&u, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
