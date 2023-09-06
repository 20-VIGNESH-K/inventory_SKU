package controllers

import (
	"net/http"

	"github.com/20-VIGNESH-K/inventory_SKU/interfaces"
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
		Quantity float32 `json:"quantity"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UpdatedInventoryService.UpdatedInventory(request.SKU, request.Quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inStock, err := c.UpdatedInventoryService.IsInStock(request.SKU)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if inStock {
		ctx.JSON(http.StatusOK, gin.H{"message": "Item is in stock"})
		ctx.JSON(http.StatusOK, gin.H{"message": "Item sold successfully"})

	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Item is not in stock"})
		return
	}

}
