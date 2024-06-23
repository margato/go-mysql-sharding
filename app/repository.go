package main

import (
	"errors"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Shards map[int]*gorm.DB
}

func NewCustomerGormRepository(shards map[int]*gorm.DB) *CustomerRepository {
	return &CustomerRepository{shards}
}

func (repository CustomerRepository) Save(customer *Customer) error {
	shard := HashShard(customer.Id, len(repository.Shards))
	db := repository.Shards[shard]

	customer.SetShard(shard)
	return db.Create(customer).Error
}

func (repository CustomerRepository) GetById(id string) (*Customer, error) {
	shard := HashShard(id, len(repository.Shards))
	db := repository.Shards[shard]

	var customer Customer
	result := db.First(&customer, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("customer not found")
	}

	customer.SetShard(shard)
	return &customer, nil
}
