package model

type Origin struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Regency struct {
	ID         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	ProvinceID string `json:"province_id" bson:"province_id"`
}

type District struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	RegencyID string `json:"regency_id" bson:"regency_id"`
}

type CustomerModel struct {
	ID          string   `bson:"id" json:"id"`
	Name        string   `bson:"name" json:"name"`
	Phone       string   `bson:"phone" json:"phone"`
	Province    Origin   `bson:"province" json:"province"`
	City        Regency  `bson:"city" json:"city"`
	District    District `bson:"district" json:"district"`
	Location    Location `bson:"location" json:"location"`
	Picture     string   `bson:"picture" json:"picture"`
	Thumbnail   string   `bson:"thumbnail" json:"thumbnail"`
	Description string   `bson:"description" json:"description"`
	CreatedAt   int64    `bson:"createdAt" json:"createdAt"`
}

type CustomerRequestPayload struct {
	Name        string   `bson:"name"`
	Phone       string   `bson:"phone"`
	Province    Origin   `bson:"province"`
	City        Regency  `bson:"city"`
	District    District `bson:"district"`
	Location    Location `bson:"location" json:"location"`
	Picture     string   `bson:"picture"`
	Thumbnail   string   `bson:"thumbnail"`
	Description string   `bson:"description"`
	CreatedAt   int64    `bson:"createdAt"`
}
