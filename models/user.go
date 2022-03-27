package models

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     string  `json:"age" bson:"age"`
	Address Address `json:"address" bson:"address"`
}

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}
