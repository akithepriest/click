package database

import (
	"context"
	"database/sql"
	"errors"
)

// Database actions related to the models.
type DatabaseModel interface {
    Create() (userID int, err error)
    Delete() (done bool, err error)
    GetByID(id int) 
    Exec(cmd string, args ...interface{}) error
}

// Represents the "users" table in database.
type User struct {
    db *PostgresDB

    ID int
    Name string 
    Email string
}

func NewUser(db *PostgresDB, name string, email string) *User {
    return &User{
        db: db,
        Name: name,
        Email: email,
    }
}

func (user *User) Create(ctx context.Context) (userID int, err error) {
    cmd := `
    INSERT INTO users (name, email)
    VALUES ($1, $2)
    RETURNING id
    `
    err = user.db.db.QueryRow(ctx, cmd, user.Name, user.Email).Scan(&userID)
    if err != nil {
        return userID, err
    } 
    return userID, nil 
}

func (user *User) Delete(ctx context.Context) (done bool, err error) {
    cmd := `
    DELETE FROM users WHERE id=$1
    `
    cmdTag, err := user.db.db.Exec(ctx, cmd, user.ID)
    if err != nil {
        return false, err
    }
    if cmdTag.RowsAffected() != 1 {
        return false, errors.New("no rows found to delete")
    }
    return true, nil
}

func (user *User) GetByID(ctx context.Context, id int) (fetchedUser *User, err error) {
    var name, email string 
    cmd := `
    SELECT name, email FROM users WHERE id=$1
    `
    err = user.db.db.QueryRow(ctx, cmd, id).Scan(&name, &email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("no user found with that ID")
        }
        return nil, err 
    }
    fetchedUser = &User{
        db: user.db,
        ID: id,
        Name: name,
        Email: email,
    }
    return fetchedUser, nil 
}