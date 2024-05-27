package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetGifts(t *testing.T) {
	expected := twitcasting.GiftContainer{
		SliceId: "1",
		Gifts: []twitcasting.Gift{
			{
				Id:             "1",
				Message:        "test_message",
				ItemImage:      "https://example.com/image.png",
				ItemSubImage:   "https://example.com/sub_image.png",
				ItemId:         "1",
				ItemMp:         "100",
				ItemName:       "test_item",
				UserImage:      "https://example.com/user_image.png",
				UserScreenId:   "test_screen_id",
				UserScreenName: "test_screen_name",
				UserName:       "test_user",
			},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	giftsResponse, errorResponse, err := locator.Gift.GetGifts()
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, giftsResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	giftsResponse, errorResponse, err = locator.Gift.GetGifts()
	assert.Nil(t, giftsResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
