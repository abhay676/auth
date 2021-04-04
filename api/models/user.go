package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// User: model for "user" collection
type User struct {
	UserID    string `json:"user_id" bson:"user_id"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Name      string `json:"name" bson:"name"`
	IP        string `json:"ip" bson:"ip"`
	UserAgent string `json:"user_agent" bson:"user_agent"`
	Token     string `json:"token" bson:"token"`
}

//Validate: used to check the params pass from Client side
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			return errors.New("password required")
		}
		if u.Email == "" {
			return errors.New("email required")
		}
	case "signup":
		if u.Name == "" {
			return errors.New("name is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
	}
	return nil
}

func (u *User) Save(db *mongo.Database) (*User, error) {
	_, err := db.Collection("user").InsertOne(context.TODO(), &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func (u *User) FindByEmail(db *mongo.Database, email string) (*User, error) {
	user := &User{}
	err := db.Collection("user").FindOne(context.TODO(), bson.D{{"email", email}}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}