package model

// Product represents a product that can be sell on the app.
type Product struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Pic          string            `json:"pic"`
	Agreement    string            `json:"agreement"`
	ActiveWeight int               `json:"active_weight"`
	Category     FireworksCategory `json:"category"`
}
