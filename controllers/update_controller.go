package controllers

import (
	"inventory_SKU/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateInventoryController struct {
	UpdatedInventoryService interfaces.IUpdateInventory
}

func InitUpdateInventoryController(UpdatedInventoryService interfaces.IUpdateInventory) UpdateInventoryController {
	return UpdateInventoryController{UpdatedInventoryService}
}

func (c *UpdateInventoryController) UpdatedInventory(ctx *gin.Context) {
	var request struct {
		SKU      string  `json:"sku"`
		Quantity float64 `json:"quantity"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UpdatedInventoryService.UpdatedInventory(request.SKU, request.Quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Item sold successfully"})
}
