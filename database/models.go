package database

// Represent a user in "users" table 
type User struct {
    ID int
    Name string
    Email string
}

// Represent a url in "urls" table
type Url struct {
    ID int
    ShortUrl string
    RedirectUrl string
    UserID int
}