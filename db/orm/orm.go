package orm

import (
	"louis/pw/db"
	"louis/pw/model"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(dbName string) *db.Store {
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = Migrate(conn)
	if err != nil {
		panic(err)
	}

	return &db.Store{
		User: &dbUser{
			conn: conn,
		},
		Client: &dbClient{
			conn: conn,
		},
		Product: &dbProduct{
			conn: conn,
		},
		Sell: &dbSell{
			conn: conn,
		},
	}
}

func NewPostgre(dsn string) *db.Store {
	dsn = "host=db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = Migrate(conn)
	if err != nil {
		panic(err)
	}

	return &db.Store{
		User: &dbUser{
			conn: conn,
		},
		Client: &dbClient{
			conn: conn,
		},
		Product: &dbProduct{
			conn: conn,
		},
		Sell: &dbSell{
			conn: conn,
		},
	}
}

func Migrate(conn *gorm.DB) error {
	err := conn.AutoMigrate(&model.User{}, &model.Client{}, &model.Product{}, &model.Sell{})
	if err != nil {
		return err
	}

	return nil
}
