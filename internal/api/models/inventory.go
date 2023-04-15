package models

type InventoryData struct {
	AssetData           []AssetData
	TotalInventoryCount int
}

type AssetData struct {
	AssetID         string
	Amount          int
	Name            string
	ClassID         string
	MarketHashName  string
	NameColor       string
	Tradable        int
	BackgroundColor string
	IconURL         string
	IconURLLarge    string
	Price           int
}
