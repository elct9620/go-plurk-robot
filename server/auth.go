package server

import (
	"github.com/elct9620/go-plurk-robot/db"
	"github.com/labstack/echo"
	"net"
	"net/http"
	"strings"
)

func AuthMiddleware() echo.HandlerFunc {
	return func(c *echo.Context) (err error) {
		mdb, err := getDatabase()
		userInfo, _ := cookie.Get(c.Request(), "SESSION")

		username, ok := userInfo.Values["username"].(string)
		password, ok := userInfo.Values["password"].(string)

		defer mdb.Session.Close()

		authorized := db.AuthorizeUser(mdb, username, password)
		ipValid := db.ValidUserIP(mdb, username, getClientIP(c.Request()))

		if (!ok || !authorized || !ipValid) && CheckProtected(c.Request().RequestURI) {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		// TODO(elct9620): This should move into login handler
		if (authorized && ipValid) && c.Request().RequestURI == "/login" {
			return c.Redirect(http.StatusFound, "/")
		}

		return
	}
}

func CheckProtected(uri string) bool {
	// NOTE(elct9620): Simple let login didn't do redirect
	// TODO(elct9620): Need to implement more ACL feature

	switch true {
	case strings.Index(uri, "/css") == 0:
		return false
	case strings.Index(uri, "/js") == 0:
		return false
	case strings.Index(uri, "/vendor") == 0:
		return false
	case strings.Index(uri, "/img") == 0:
		return false
	case (uri == "/login"):
		return false
	}

	return true
}

// Try get user IP
// TODO(elct9620): This function should use more security way
func getClientIP(r *http.Request) string {

	// Fetch All Available IP
	clientIP := r.Header.Get("HTTP_CLIENT_IP")
	xForwardedFor := r.Header.Get("HTTP_X_FORWARDED_FOR")
	xForwarded := r.Header.Get("HTTP_X_FORWARDED")
	remoteAddr := r.Header.Get("REMOTE_ADDR")

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	ipList := []string{clientIP, xForwardedFor, xForwarded, remoteAddr}

	if len(ip) <= 0 {
		for _, v := range ipList {
			if len(v) > 0 {
				ip = v
				break
			}
		}
	}

	return ip
}

func login(c *echo.Context) (err error) {
	return c.Render(200, "login", nil)
}

func verifyLogin(c *echo.Context) (err error) {

	mdb, err := getDatabase()
	defer mdb.Session.Close()

	username := c.Request().FormValue("username")
	password := c.Request().FormValue("password")

	encryptedPassword := db.EncryptPassword(password)

	if db.AuthorizeUser(mdb, username, encryptedPassword) {
		userInfo, _ := cookie.Get(c.Request(), "SESSION")
		userInfo.Values["username"] = username
		userInfo.Values["password"] = encryptedPassword

		err = userInfo.Save(c.Request(), c.Response())
		err = db.RefreshUserIP(mdb, &db.User{username, encryptedPassword, ""}, getClientIP(c.Request()))
	} else {
		// TODO(elct9620): Show error message to notice user
	}

	return c.Redirect(http.StatusFound, "/")
}
