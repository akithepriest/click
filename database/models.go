package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represent user document
type User struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
}

// Represent url document.
type Url struct {
	ID primitive.ObjectID `bson:"_id"`
	Vanity string `bson:"vanity"`
	RedirectUrl string `bson:"redirect_url"`
	UserID primitive.ObjectID `bson:"user_id"`
}

// Represent click document.
// A click is stored when a vanity url is visited. 
type Click struct {
	ID primitive.ObjectID `bson:"_id"`
	UrlID primitive.ObjectID `bson:"url_id"`
	UrlUserID primitive.ObjectID `bson:"url_user_id"`
	ClickedOn primitive.Timestamp `bson:"clicked_on"`
	Agent string `bson:"agent"`
}