package db

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpenSession(t *testing.T) {
	session, _ := OpenSession("mongodb:///dev-test")

	// Test database connect working, and open specify database
	assert.Equal(t, "dev-test", session.DB("").Name)
}

func Test_OpenSessionUsingEnv(t *testing.T) {
	os.Setenv("MONGODB_URL", "MONGO_URL")
	os.Setenv("MONGO_URL", "mongodb:///dev-test")
	session, _ := OpenSession("")

	assert.Equal(t, "dev-test", session.DB("").Name)
}
