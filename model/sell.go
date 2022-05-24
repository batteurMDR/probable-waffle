package model

// Sell represents a client buying some products.
type Sell struct {
	ID            string `json:"id"`
	Date          string `json:"date"`
	Reason        string `json:"reason"`
	Products      string `json:"products"`
	ClientID      string `json:"client_id"`
	PaymentMethod string `json:"payment_method"`
}
