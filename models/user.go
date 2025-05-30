package models

type User struct {
	UserID      int64
	Username    string
	Password    string
	PhoneNumber string
	Email       string
	FirstName   string
	MiddleName  string
	LastName    string
}

type Address struct {
	AddressID   int64
	AddressName string
	Country     string
	State       string
	PinCode     string
	LandMark    string
}
