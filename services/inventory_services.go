package services

import (
	"context"
	"fmt"
	"inventory_SKU/interfaces"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	InventoryCollection *mongo.Collection
	ctx                 context.Context
}

func NewInventoryServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IInventory {
	return &CustomerService{collection, ctx}

}
func (b *CustomerService) CreateInventory(customers []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := b.InventoryCollection.InsertMany(ctx, customers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil

}
