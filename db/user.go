package db

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string
	Password string
	LastIP   string // For secure the user login
}

func CreateUser(db *mgo.Database, username string, password string) (user *User, err error) {
	users := db.C("users")
	count, err := users.Find(bson.M{"username": username}).Count()
	if err != nil {
		return
	}

	if count > 0 {
		return nil, errors.New("User exists!")
	}

	user = &User{Username: username, Password: EncryptPassword(password)}
	err = users.Insert(*user)

	return
}

func AuthorizeUser(db *mgo.Database, username string, password string) bool {
	users := db.C("users")
	count, err := users.Find(bson.M{"username": username, "password": password}).Count()

	if err != nil {
		return false
	}

	if count != 1 {
		return false
	}

	return true
}

// For security, the server should check user IP is same as last login
func ValidUserIP(db *mgo.Database, username string, ip string) bool {
	users := db.C("users")
	if count, err := users.Find(bson.M{"username": username, "lastip": ip}).Count(); err != nil || count != 1 {
		return false
	}
	return true
}

func RefreshUserIP(db *mgo.Database, user *User, ip string) error {
	users := db.C("users")
	user.LastIP = ip
	return users.Update(bson.M{"username": user.Username}, user)
}

func EncryptPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
