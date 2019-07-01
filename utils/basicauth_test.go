package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
)

func fakeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func makeRequest(cfg Config, username string, password string, t *testing.T) (ret *httptest.ResponseRecorder) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth(username, password)
	if err != nil {
	}

	client.Do(req)

	rr := httptest.NewRecorder()
	http.HandlerFunc(BasicAuth(fakeHandler, cfg)).ServeHTTP(rr, req)
	return rr
}

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Basic auth func", func() {

		cfg := Config{"test", "username", "password"}

		g.Describe("When no basic auth is given", func() {
			g.It("returns Unauthorized", func() {
				rr := makeRequest(cfg, "", "", t)
				status := rr.Code
				g.Assert(status).Equal(http.StatusUnauthorized)
			})
		})

		g.Describe("When no basic auth is wrong", func() {
			g.It("returns Unauthorized", func() {
				rr := makeRequest(cfg, "wrong", "stuff", t)
				status := rr.Code
				g.Assert(status).Equal(http.StatusUnauthorized)
			})
		})

		g.Describe("When basic auth is good", func() {
			g.It("returns Unauthorized", func() {
				rr := makeRequest(cfg, "username", "password", t)
				status := rr.Code
				g.Assert(status).Equal(http.StatusOK)
			})
		})
	})
}
