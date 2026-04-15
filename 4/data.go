package main

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

type ProductModel struct {
	Id          uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	Name        string
	Description string
	Price       uint32
}

type Cart = []CartItem

type CartItem struct {
	Id uuid.UUID `json:"id"`
	Product
}

type CartModel struct {
	Id       uuid.UUID      `gorm:"primaryKey;autoIncrement:false"`
	Products []ProductModel `gorm:"many2many:cart_products;"`
}

type Database struct {
	ctx context.Context
	db  *gorm.DB
}

func NewDatabase(path string) (Database, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	db.AutoMigrate(&ProductModel{})
	db.AutoMigrate(&CartModel{})

	return Database{context.Background(), db}, nil
}

func (p *Database) ExistsProduct(id uuid.UUID) (bool, error) {
	models, err := gorm.G[ProductModel](p.db).Where("id = ?", id).Find(p.ctx)
	if err != nil {
		return false, err
	}

	return len(models) > 0, nil
}

func (p *Database) ListProduct() (map[uuid.UUID]Product, error) {
	models, err := gorm.G[ProductModel](p.db).Find(p.ctx)
	if err != nil {
		return nil, err
	}

	products := make(map[uuid.UUID]Product)
	for _, model := range models {
		product := Product{
			Name:        model.Name,
			Description: model.Description,
			Price:       model.Price,
		}

		products[model.Id] = product
	}

	return products, nil
}

func (p *Database) CreateProduct(product Product) (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return id, err
	}

	model := ProductModel{
		Id:          id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	err = gorm.G[ProductModel](p.db).Create(p.ctx, &model)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (p *Database) ReadProduct(id uuid.UUID) (Product, error) {
	model, err := gorm.G[ProductModel](p.db).Where("id = ?", id).First(p.ctx)
	if err != nil {
		return Product{}, err
	}

	product := Product{
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
	}

	return product, nil
}

func (p *Database) UpdateProduct(id uuid.UUID, product Product) error {
	model := ProductModel{
		Id:          id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	_, err := gorm.G[ProductModel](p.db).Where("id = ?", model.Id).Updates(p.ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (p *Database) DeleteProduct(id uuid.UUID) error {
	_, err := gorm.G[ProductModel](p.db).Where("id = ?", id).Delete(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Database) ExistsCart(id uuid.UUID) (bool, error) {
	models, err := gorm.G[CartModel](p.db).Where("id = ?", id).Find(p.ctx)
	if err != nil {
		return false, err
	}

	return len(models) > 0, nil
}

func (p *Database) CreateCart() (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return id, err
	}

	model := CartModel{
		Id:       id,
		Products: make([]ProductModel, 0),
	}

	err = gorm.G[CartModel](p.db).Create(p.ctx, &model)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (p *Database) ReadCart(id uuid.UUID) (Cart, error) {
	model, err := gorm.G[CartModel](p.db).Where("id = ?", id).First(p.ctx)
	if err != nil {
		return nil, err
	}

	products := make([]ProductModel, 0)
	err = p.db.Model(&model).Association("Products").Find(&products)
	if err != nil {
		return nil, err
	}

	cart := make(Cart, 0)
	for _, product := range products {
		cart = append(cart, CartItem{
			Id: product.Id,
			Product: Product{
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
			},
		})
	}

	return cart, nil
}

func (p *Database) DeleteCart(id uuid.UUID) error {
	_, err := gorm.G[CartModel](p.db).Where("id = ?", id).Delete(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Database) AddToCart(cid uuid.UUID, pid uuid.UUID) error {
	cart, err := gorm.G[CartModel](p.db).Where("id = ?", cid).First(p.ctx)
	if err != nil {
		return err
	}

	product, err := gorm.G[ProductModel](p.db).Where("id = ?", pid).First(p.ctx)
	if err != nil {
		return err
	}

	err = p.db.Model(&cart).Association("Products").Append(product)
	if err != nil {
		return err
	}

	return nil
}
