package model

type UserModel struct {
	ID       int    `bson:"id" json:"id,omitempty"`
	Name     string `bson:"name" json:"name,omitempty"`
	Username string `bson:"username" json:"username,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
	Role     string `bson:"role" json:"role,omitempty"`
}
