package models

type Report struct {
	Name     string `bson:"name"`
	Address  string `bson:"address"`
	City     string `bson:"city"`
	Country  string `bson:"country"`
	Pincode  string `bson:"pincode"`
	SATScore int    `bson:"score"`
	Passed   bool   `bson:"passed"`
}
