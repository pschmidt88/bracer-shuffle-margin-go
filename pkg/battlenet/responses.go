package battlenet

type ConnectedRealm struct {
	ID int
}

type AuctionListResponse struct {
	Auctions []Auction
}

type Item struct {
	ID int
}

type Auction struct {
	Buyout   int
	Item     Item
	Quantity int
}
