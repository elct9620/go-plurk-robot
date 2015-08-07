package server

import (
	"github.com/gorilla/sessions"
	"os"
)

func NewCookieStore() *sessions.CookieStore {
	cookieName := os.Getenv("SECRET_KEY")
	if len(cookieName) <= 0 {
		cookieName = "SIMPLE_COOKIE_KEY"
	}

	cookie := sessions.NewCookieStore([]byte(cookieName))
	cookie.Options.HttpOnly = true

	return cookie
}
