package model

// Client represents a customer who can buy products.
type Client struct {
	ID              string   `json:"id"`
	Valid           bool     `json:"valid"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	BirthDate       string   `json:"birth_date"`
	BirthPlace      string   `json:"birth_place"`
	IdCardPath      string   `json:"id_card_path"`
	IdCardType      CardType `json:"id_card_type"`
	IdCardNum       string   `json:"id_card_num"`
	IdCardAuthority string   `json:"id_card_authority"`
	IdCardDate      string   `json:"id_card_date"`
}
