package main

import (
	"context"
	"fmt"
	"inventory_SKU/config"
	"inventory_SKU/constants"
	"inventory_SKU/controllers"
	"inventory_SKU/services"
	"log"

	routes "inventory_SKU/route"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	//mongoclient *mongo.Client
	ctx    context.Context
	server *gin.Engine
)

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	inventoryCollection := mongoClient.Database(constants.DatabaseName).Collection("items")
	inventoryService := services.NewInventoryServiceInit(inventoryCollection, ctx)
	inventoryController := controllers.InitInventoryController(inventoryService)
	routes.Route(server, inventoryController)
}
func initUpdateApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	inventoryCollection := mongoClient.Database(constants.DatabaseName).Collection("items")
	inventoryService := services.NewUpdatedInventoryServiceInit(inventoryCollection)
	inventoryController := controllers.InitUpdateInventoryController(inventoryService)
	routes.UpdateRoute(server, inventoryController)

}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	initApp(mongoclient)
	initUpdateApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
