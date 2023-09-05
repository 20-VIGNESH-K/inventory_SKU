package routes

import (
	"inventory_SKU/controllers"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, controller controllers.InventoryController) {
	router.POST("/api/inventory/create", controller.CreateInventory)

}

func UpdateRoute(router *gin.Engine, controller controllers.UpdateInventoryController) {
	router.POST("/api/inventory/update", controller.UpdatedInventory)
}
