package main

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Customer struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Shard     int       `json:"shard" gorm:"-"`
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Id:        ulid.Make().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Shard:     -1,
	}
}

func (c *Customer) SetShard(shard int) {
	c.Shard = shard
}
