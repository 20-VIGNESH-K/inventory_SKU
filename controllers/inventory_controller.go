package controllers

import (
	"net/http"
	"strings"

	"github.com/20-VIGNESH-K/inventory_SKU/interfaces"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryService interfaces.IInventory
}

func InitInventoryController(inventoryService interfaces.IInventory) InventoryController {
	return InventoryController{inventoryService}
}

func (tc *InventoryController) CreateInventory(ctx *gin.Context) {
	var inventory []interface{}
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newInventory, err := tc.InventoryService.CreateInventory(inventory)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newInventory})
}
