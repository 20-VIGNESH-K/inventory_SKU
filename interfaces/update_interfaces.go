package interfaces

type IUpdateInventory interface {
	UpdatedInventory(sku string, soldQuantity float64) error
	IsInStock(sku string) (bool, error)
}
