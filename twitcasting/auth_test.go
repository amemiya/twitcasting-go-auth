package twitcasting_test

import (
	"encoding/json"
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthorizeUrl(t *testing.T) {
	locator, _ := twitcasting.CreateServiceLocator(nil, &twitcasting.BasicLogger{Logger: &log.Logger{}}, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	url := locator.Auth.GetAuthorizeUrl("client", "state")
	assert.Equal(t, "https://apiv2.twitcasting.tv/oauth2/authorize?client_id=client&response_type=code&state=state", url)
}

func TestPostAccessToken(t *testing.T) {
	expected := twitcasting.AccessTokenContainer{
		TokenType:   "bearer",
		ExpiresIn:   86400,
		AccessToken: "access_token",
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(
			t,
			http.Header{
				"Accept-Encoding": {"gzip"},
				"Content-Length":  {"124"},
				"Content-Type":    {"application/x-www-form-urlencoded"},
				"User-Agent":      {"Go-http-client/1.1"},
			},
			r.Header,
		)
		body, _ := json.Marshal(expected)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-We-Are-Hiring", "https://about.moi.st/ja/recruit/")
		w.WriteHeader(http.StatusOK)
		writeString, _ := io.WriteString(w, string(body))
		assert.Equal(t, len(body), writeString)
	})
	listener, _ := net.Listen("tcp", ":8080")
	server := httptest.Server{Listener: listener, Config: &http.Server{Handler: handler}}

	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	accessToken, errorResponse, err := locator.Auth.PostAccessToken("client", "secret", "code", "http://localhost/callback")
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, accessToken)
	server.Close()
}
