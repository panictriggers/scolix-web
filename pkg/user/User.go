package user

type User struct {
	DatabaseID 	string `bson:"_id" json:"_"` // For MongoDB, but we are not sure if we will use MongoDB
	Username	string `bson:"username" json:"username"`
}
