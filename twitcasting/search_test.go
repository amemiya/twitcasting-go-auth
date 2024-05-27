package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSearchUsers(t *testing.T) {
	expected := twitcasting.SearchUsersContainer{
		Users: []twitcasting.User{
			{
				Id:              "456543234",
				Name:            "test_name",
				ScreenId:        "screen_id",
				Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
				Profile:         "test_profile",
				Level:           30,
				LatestMovieId:   "123432345",
				IsLive:          true,
				SupporterCount:  0,
				SupportingCount: 0,
				Created:         0,
			},
		},
	}

	server := CreateTestSever(t, expected, "", url.Values{"lang": []string{"ja"}, "limit": []string{"10"}, "words": []string{"test"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	usersResponse, errorResponse, err := locator.Search.SearchUsers("test", 10, false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, usersResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"lang": []string{"ja"}, "limit": []string{"10"}, "words": []string{"test"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	usersResponse, errorResponse, err = locator.Search.SearchUsers("test", 10, true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, usersResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"lang": []string{"ja"}, "limit": []string{"10"}, "words": []string{"test"}}, false, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	usersResponse, errorResponse, err = locator.Search.SearchUsers("test", 10, false)
	assert.Nil(t, usersResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestSearchLiveMovies(t *testing.T) {
	expected := twitcasting.SearchLiveMoviesContainer{
		Movies: []twitcasting.MovieContainer{
			{
				Movie: twitcasting.Movie{
					Id:               "123432345",
					UserId:           "456543234",
					Title:            "test_title",
					Subtitle:         "test_subtitle",
					LastOwnerComment: "test_last_owner_comment",
					Category:         "test_category",
					Link:             "https://twitcasting.tv/test_owner_id/movie/123432345",
					IsLive:           true,
					IsRecorded:       false,
					CommentCount:     0,
					LargeThumbnail:   "https://twitcasting.tv/usr/test_owner_id/thumb/medium",
					SmallThumbnail:   "https://twitcasting.tv/usr/test_owner_id/thumb/small",
					Country:          "ja",
					Duration:         0,
					Created:          0,
					IsCollabo:        false,
					IsProtected:      false,
					MaxViewCount:     0,
					CurrentViewCount: 0,
					TotalViewCount:   0,
					HlsUrl:           "https://twitcasting.tv/usr/test_owner_id/movie/123432345/hls",
				},
				Broadcaster: twitcasting.Broadcaster{
					Id:              "456543234",
					ScreenId:        "screen_id",
					Name:            "test_name",
					Image:           "https://img.twitcasting.tv/usr/test_id/thumb/medium",
					Profile:         "test_profile",
					Level:           30,
					LastMovieId:     "123432345",
					IsLive:          true,
					SupporterCount:  0,
					SupportingCount: 0,
					Created:         0,
				},
				Tags: []string{"test_tag", "test_tag2"},
			},
		},
	}

	server := CreateTestSever(t, expected, "", url.Values{"context": []string{"test_word"}, "lang": []string{"ja"}, "limit": []string{"10"}, "type": []string{"word"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err := locator.Search.SearchLiveMovies("word", "test_word", 10, false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, moviesResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"context": []string{"test_word"}, "lang": []string{"ja"}, "limit": []string{"10"}, "type": []string{"word"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err = locator.Search.SearchLiveMovies("word", "test_word", 10, true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, moviesResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"context": []string{"test_word"}, "lang": []string{"ja"}, "limit": []string{"10"}, "type": []string{"word"}}, false, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err = locator.Search.SearchLiveMovies("word", "test_word", 10, false)
	assert.Nil(t, moviesResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
