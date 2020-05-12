package tsm

// Item represents an Item from the auction house in the TSM API
type Item struct {
	ID          int
	Name        string
	VendorBuy   int
	VendorSell  int
	MarketValue int
	MinBuyout   int
}
