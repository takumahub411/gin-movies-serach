package initializers

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/gorilla/sessions"
	"io"
	"strings"
)

var SessionName = "gsid"

var Store *sessions.CookieStore

var Session *sessions.Session

func SetSessionOption() {
	Store.Options.HttpOnly = true
	Store.Options.Secure = false
	Store.Options.Domain = "localhost"
	Store.Options.MaxAge = 0
	//gob.Register(&models.User{})
}

func SessionInit() {
	b := make([]byte, 48)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}

	sessionKey := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	Store = sessions.NewCookieStore([]byte(sessionKey))
	Session = sessions.NewSession(Store, SessionName)

	SetSessionOption()

}
