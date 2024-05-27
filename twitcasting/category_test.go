package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetCategories(t *testing.T) {
	expected := twitcasting.CategoriesContainer{
		Categories: []twitcasting.Category{
			{
				Id:   "1",
				Name: "test_category",
				SubCategories: []twitcasting.SubCategory{
					{
						Id:    "1",
						Name:  "test_sub_category",
						Count: 10,
					},
				},
			},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"lang": []string{"ja"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	categoriesResponse, errorResponse, err := locator.Category.GetCategories("ja", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, categoriesResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"lang": []string{"en"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	categoriesResponse, errorResponse, err = locator.Category.GetCategories("en", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, categoriesResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "", url.Values{"lang": []string{"ja"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	categoriesResponse, errorResponse, err = locator.Category.GetCategories("ja", true)
	assert.Nil(t, categoriesResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
