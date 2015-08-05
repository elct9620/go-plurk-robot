// Database helper
package db

import (
	"gopkg.in/mgo.v2"
	"os"
	"strings"
)

func OpenSession(server string) (session *mgo.Session, err error) {
	// If no server specify, use environment value
	if len(server) <= 0 {
		server = os.Getenv("MONGODB_URL")

		// If environment variable reference to another variable
		if strings.Index(server, "mongodb:") != 0 {
			server = os.Getenv(server)
		}
	}

	session, err = mgo.Dial(server)
	return session, err
}
