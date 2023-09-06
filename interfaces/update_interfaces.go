package interfaces

type IUpdateInventory interface {
	UpdatedInventory(sku string, soldQuantity float32) error
	IsInStock(sku string) (bool, error)
}
