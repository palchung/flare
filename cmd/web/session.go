package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager

func NewSession() *scs.SessionManager {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProduction

	return session
}
