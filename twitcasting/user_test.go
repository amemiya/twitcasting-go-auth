package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetUser(t *testing.T) {
	expected := twitcasting.UserContainer{
		User: twitcasting.User{
			Id:              "test_id",
			Name:            "test_name",
			ScreenId:        "screen_id",
			Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
			Profile:         "test_profile",
			Level:           30,
			LatestMovieId:   "123432345",
			IsLive:          true,
			SupporterCount:  0,
			SupportingCount: 0,
			Created:         1716666000,
		},
		SupporterCount:  2000,
		SupportingCount: 100,
	}
	server := CreateTestSever(t, expected, "", url.Values{}, false, http.StatusOK)

	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	userResponse, errorResponse, err := locator.User.GetUser("test", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, userResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator2 := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	userResponse, errorResponse, err = locator2.User.GetUser("test", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, userResponse)
	server.Close()

	expected3 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected3, "", url.Values{}, false, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	userResponse, errorResponse, err = locator.User.GetUser("test", false)
	assert.Nil(t, userResponse)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected3, errorResponse)
	server.Close()
}

func TestGetVerifyCredentials(t *testing.T) {
	expected := twitcasting.VerifyCredentialsContainer{
		App: twitcasting.App{
			ClientId:    "test_client_id",
			Name:        "test_name",
			OwnerUserId: "owner_user_id",
		},
		User: twitcasting.User{
			Id:              "test_id",
			Name:            "test_name",
			ScreenId:        "screen_id",
			Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
			Profile:         "test_profile",
			Level:           30,
			LatestMovieId:   "123432345",
			IsLive:          true,
			SupporterCount:  0,
			SupportingCount: 0,
			Created:         1716666000,
		},
		SupporterCount:  2000,
		SupportingCount: 100,
	}
	server := CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	userResponse, errorResponse, err := locator.User.GetVerifyCredentials()
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, userResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	userResponse, errorResponse, err = locator.User.GetVerifyCredentials()
	assert.Nil(t, userResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
