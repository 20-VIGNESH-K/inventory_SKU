package models

type Inventory_SKU struct {
	Sku      string       `json:"sku" bson:"sku"`
	Price    Price_type   `json:"price" bson:"price"`
	Quantity float32      `json:"quantity" bson:"quantity,truncate"`
	Options  Options_type `json:"options" bson:"options"`
}

type Price_type struct {
	Base     float32 `json:"base" bson:"base,truncate"`
	Currency string  `json:"currency" bson:"currency"`
	Discount float32 `json:"discount" bson:"discount,truncate"`
}
type Options_type struct {
	Size     Size_type `json:"size" bson:"size"`
	Features []string  `json:"features" bson:"features"`
	Colors   []string  `json:"colors" bson:"colors"`
	Ruling   string    `json:"ruling" bson:"ruling"`
	Image    string    `json:"image" bson:"image"`
}

type Size_type struct {
	H float32 `json:"h" bson:"h,truncate"`
	L float32 `json:"l" bson:"l,truncate"`
	W float32 `json:"w" bson:"w,truncate"`
}
