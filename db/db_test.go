package db

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpenSession(t *testing.T) {
	db := os.Getenv("MONGODB_URL")
	if len(db) <= 0 {
		db = "mongodb:///"
	}
	session, _ := OpenSession(db)

	// Test database connect working, and open specify database
	assert.Equal(t, "test", session.DB("").Name)
}

func Test_OpenSessionUsingEnv(t *testing.T) {
	db := os.Getenv("MONGODB_URL")
	if len(db) <= 0 {
		os.Setenv("MONGODB_URL", "MONGO_URL")
		os.Setenv("MONGO_URL", "mongodb:///test")
	}
	session, _ := OpenSession("")

	assert.Equal(t, "test", session.DB("").Name)
}
