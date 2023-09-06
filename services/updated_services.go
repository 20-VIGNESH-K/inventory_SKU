package services

import (
	"context"
	"errors"
	"fmt"
	"inventory_SKU/interfaces"
	"inventory_SKU/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryService struct {
	collection *mongo.Collection
}

func NewUpdatedInventoryServiceInit(collection *mongo.Collection) interfaces.IUpdateInventory {
	return &InventoryService{
		collection: collection,
	}

}

func (s *InventoryService) IsInStock(sku string) (bool, error) {
	filter := bson.M{"sku": sku}
	var item models.Inventory_SKU

	err := s.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Item not found, consider it out of stock
		}
		return false, err
	}

	return item.Quantity >= 0, nil // In stock if quantity is greater than zero
}

func (s *InventoryService) UpdatedInventory(sku string, soldQuantity float64) error {
	filter := bson.M{"sku": sku}
	var item models.Inventory_SKU
	err := s.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("item with SKU %s not found", sku)
		}
		return err
	}

	if item.Quantity < soldQuantity {
		return fmt.Errorf("not enough quantity in stock, There is only %v available", item.Quantity)
	}
	update := bson.M{"$inc": bson.M{"quantity": -soldQuantity}}

	_, err1 := s.collection.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		return err1
	}

	return nil
}
