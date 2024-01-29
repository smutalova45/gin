package storage

import "main.go/api/models"

type IStorage interface {
	Close()
	User() IUserStorage
	Category() ICategoryStorage
	Product() IProductStorage
	Basket() IBasketStorage
	BasketProduct() IBasketProductStorage
}

type IUserStorage interface {
	Create(models.CreateUser) (string, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (string, error)
	Delete(models.PrimaryKey) error
	GetPassword(id string) (string, error)
	UpdatePassword(password models.UpdateUserPassword) error
}

type ICategoryStorage interface {
	Create(models.CreateCategory) (string, error)
	GetByID(models.PrimaryKey) (models.Category, error)
	GetList(models.GetListRequest) (models.CategoryResponse, error)
	Update(models.UpdateCategory) (string, error)
	Delete(models.PrimaryKey) error
}

type IProductStorage interface {
	Create(models.CreateProduct) (string, error)
	GetByID(models.PrimaryKey) (models.Product, error)
	GetList(models.GetListRequest) (models.ProductResponse, error)
	Update(models.UpdateProduct) (string, error)
	Delete(models.PrimaryKey) error
	Search(map[string]int) (map[string]int, map[string]int, error)
	TakeProducts(map[string]int) error
}
type IBasketStorage interface {
	Create(models.CreateBasket) (string, error)
	GetByID(models.PrimaryKey) (models.Basket, error)
	GetList(models.GetListRequest) (models.BasketResponse, error)
	Update(models.UpdateBasket) (string, error)
	Delete(key models.PrimaryKey) error
}

type IBasketProductStorage interface {
	Create(models.CreateBasketProduct) (string, error)
	GetByID(models.PrimaryKey) (models.BasketProduct, error)
	GetList(models.GetListRequest) (models.BasketProductResponse, error)
	Update(models.UpdateBasketProduct) (string, error)
	Delete(models.PrimaryKey) error
}
