package model

type UserModel struct {
	ID       int    `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role" json:"role"`
}
