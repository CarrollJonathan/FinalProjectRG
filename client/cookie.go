package client

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func GetClientWithCookie(token string, cookies ...*http.Cookie) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	cookies = append(cookies, &http.Cookie{
		Name:  "session_token",
		Value: token,
	})

	var CookieHost string = os.Getenv("RAILWAY_STATIC_URL")
	CookieScheme := "https"
	if CookieHost == "" {
		CookieHost = "http"
		CookieHost = "localhost"
	}

	jar.SetCookies(&url.URL{
		Scheme: CookieScheme,
		Host:   CookieHost,
	}, cookies)

	c := &http.Client{
		Jar: jar,
	}

	return c, nil
}
