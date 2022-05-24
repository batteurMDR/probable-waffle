package model

type CardType int

// Represent all type of identity cards
const (
	IdCard CardType = iota
	Passport
	DriveLicence
	HuntingLicence
)
