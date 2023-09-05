package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IInventory interface {
	CreateInventory(customers []interface{}) (*mongo.InsertManyResult, error)
	
}
