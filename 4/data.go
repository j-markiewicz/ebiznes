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
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       uint32    `json:"price"`
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

	return Database{context.Background(), db}, nil
}

func (p *Database) Exists(id uuid.UUID) (bool, error) {
	models, err := gorm.G[ProductModel](p.db).Where("id = ?", id).Find(p.ctx)
	if err != nil {
		return false, err
	}

	return len(models) > 0, nil
}

func (p *Database) List() (map[uuid.UUID]Product, error) {
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

func (p *Database) Create(product Product) (uuid.UUID, error) {
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

func (p *Database) Read(id uuid.UUID) (Product, error) {
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

func (p *Database) Update(id uuid.UUID, product Product) error {
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

func (p *Database) Delete(id uuid.UUID) error {
	_, err := gorm.G[ProductModel](p.db).Where("id = ?", id).Delete(p.ctx)
	if err != nil {
		return err
	}

	return nil
}
