package db

import "louis/pw/model"

type Store struct {
	User    StoreUser
	Client  StoreClient
	Product StoreProduct
	Sell    StoreSell
}

type StoreUser interface {
	CreateUser(u *model.User) (*model.User, error)
	GetAllUser() ([]*model.User, error)
	GetUserById(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(id string, data map[string]interface{}) (*model.User, error)
	DeleteUser(id string) error
}

type StoreClient interface {
	CreateClient(u *model.Client) (*model.Client, error)
	GetAllClient() ([]*model.Client, error)
	GetClientById(id string) (*model.Client, error)
	UpdateClient(id string, data map[string]interface{}) (*model.Client, error)
	ValidateClient(id string) (*model.Client, error)
	UnvalidateClient(id string) (*model.Client, error)
	DeleteClient(id string) error
}

type StoreProduct interface {
	CreateProduct(u *model.Product) (*model.Product, error)
	GetAllProduct() ([]*model.Product, error)
	GetProductById(id string) (*model.Product, error)
	UpdateProduct(id string, data map[string]interface{}) (*model.Product, error)
	DeleteProduct(id string) error
}

type StoreSell interface {
	GetAllSell() ([]*model.Sell, error)
	GetSellById(id string) (*model.Sell, error)
	CreateSell(u *model.Sell) (*model.Sell, error)
}
