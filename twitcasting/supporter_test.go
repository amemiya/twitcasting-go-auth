package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSupportingStatus(t *testing.T) {
	expected := twitcasting.SupportingStatusContainer{
		IsSupporting: true,
		Supported:    345654323,
		TargetUser: twitcasting.User{
			Id:              "3456543",
			ScreenId:        "screen_id",
			Name:            "test_name",
			Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
			Profile:         "test_profile",
			Level:           30,
			LatestMovieId:   "123432345",
			IsLive:          true,
			SupporterCount:  0,
			SupportingCount: 0,
			Created:         0,
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"target_user_id": []string{"3456543"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supportingStatusResponse, errorResponse, err := locator.Supporter.GetSupportingStatus("12345432", "3456543", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, supportingStatusResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"target_user_id": []string{"3456543"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supportingStatusResponse, errorResponse, err = locator.Supporter.GetSupportingStatus("12345432", "3456543", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, supportingStatusResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"target_user_id": []string{"3456543"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supportingStatusResponse, errorResponse, err = locator.Supporter.GetSupportingStatus("12345432", "3456543", true)
	assert.Nil(t, supportingStatusResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestPostSupport(t *testing.T) {
	expected := twitcasting.PostSupportContainer{
		AddedCount: 2,
	}
	server := CreateTestSever(t, expected, "{\"target_user_ids\":[\"test_id\",\"test_id2\"]}", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	postSupportResponse, errorResponse, err := locator.Supporter.PostSupport([]string{"test_id", "test_id2"})
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, postSupportResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "{\"target_user_ids\":[\"test_id\",\"test_id2\"]}", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	postSupportResponse, errorResponse, err = locator.Supporter.PostSupport([]string{"test_id", "test_id2"})
	assert.Nil(t, postSupportResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestDeleteSupport(t *testing.T) {
	expected := twitcasting.DeleteSupportContainer{
		RemovedCount: 2,
	}
	server := CreateTestSever(t, expected, "{\"target_user_ids\":[\"test_id\",\"test_id2\"]}", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	deleteSupportResponse, errorResponse, err := locator.Supporter.DeleteSupport([]string{"test_id", "test_id2"})
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, deleteSupportResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "{\"target_user_ids\":[\"test_id\",\"test_id2\"]}", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	deleteSupportResponse, errorResponse, err = locator.Supporter.DeleteSupport([]string{"test_id", "test_id2"})
	assert.Nil(t, deleteSupportResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestGetSupportingList(t *testing.T) {
	expected := twitcasting.SupporterListContainer{
		Total: 1000,
		Supporting: []twitcasting.SupporterUser{
			{
				Id:              "3456543",
				Name:            "test_name",
				ScreenId:        "screen_id",
				Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
				Profile:         "test_profile",
				Level:           30,
				LastMovieId:     "123432345",
				IsLive:          true,
				Supported:       345654323,
				SupporterCount:  0,
				SupportingCount: 0,
				Created:         0,
				Point:           0,
				TotalPoint:      0,
			},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supportingListResponse, errorResponse, err := locator.Supporter.GetSupportingList("3456543", 10, 0, false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, supportingListResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supportingListResponse, errorResponse, err = locator.Supporter.GetSupportingList("3456543", 10, 0, true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, supportingListResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supportingListResponse, errorResponse, err = locator.Supporter.GetSupportingList("3456543", 10, 0, true)
	assert.Nil(t, supportingListResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestGetSupporterList(t *testing.T) {
	expected := twitcasting.SupporterListContainer{
		Total: 1000,
		Supporting: []twitcasting.SupporterUser{
			{
				Id:              "3456543",
				Name:            "test_name",
				ScreenId:        "screen_id",
				Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
				Profile:         "test_profile",
				Level:           30,
				LastMovieId:     "123432345",
				IsLive:          true,
				Supported:       345654323,
				SupporterCount:  0,
				SupportingCount: 0,
				Created:         0,
				Point:           0,
				TotalPoint:      0,
			},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}, "sort": []string{"new"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supporterListResponse, errorResponse, err := locator.Supporter.GetSupporterList("3456543", 10, 0, "new", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, supporterListResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}, "sort": []string{"new"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supporterListResponse, errorResponse, err = locator.Supporter.GetSupporterList("3456543", 10, 0, "new", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, supporterListResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}, "sort": []string{"new"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	supporterListResponse, errorResponse, err = locator.Supporter.GetSupporterList("3456543", 10, 0, "new", true)
	assert.Nil(t, supporterListResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
