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
	ID          int      `bson:"id"`
	Name        string   `bson:"name"`
	Phone       string   `bson:"phone"`
	Province    Origin   `bson:"province"`
	City        Regency  `bson:"city"`
	District    District `bson:"district"`
	Longitude   float64  `bson:"longitude"`
	Latitude    float64  `bson:"latitude"`
	Picture     string   `bson:"picture"`
	Thumbnail   string   `bson:"thumbnail"`
	Description string   `bson:"description"`
	CreatedAt   int64    `bson:"createdAt"`
}

type CustomerRequestPayload struct {
	Name        string   `bson:"name"`
	Phone       string   `bson:"phone"`
	Province    Origin   `bson:"province"`
	City        Regency  `bson:"city"`
	District    District `bson:"district"`
	Longitude   float64  `bson:"longitude"`
	Latitude    float64  `bson:"latitude"`
	Picture     string   `bson:"picture"`
	Thumbnail   string   `bson:"thumbnail"`
	Description string   `bson:"description"`
	CreatedAt   int64    `bson:"createdAt"`
}
