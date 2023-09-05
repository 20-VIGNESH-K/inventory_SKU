package services

import (
	"context"
	"inventory_SKU/interfaces"

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

func (s *InventoryService) UpdatedInventory(sku string, soldQuantity float64) error {
	filter := bson.M{"sku": sku}
	update := bson.M{"$inc": bson.M{"quantity": -soldQuantity}}

	_, err := s.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
