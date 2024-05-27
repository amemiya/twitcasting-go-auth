package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetMovie(t *testing.T) {
	expected := twitcasting.MovieContainer{
		Movie: twitcasting.Movie{
			Id:               "12345432",
			UserId:           "123456",
			Title:            "test_title",
			Category:         "test_category",
			Link:             "https://example.com/link",
			IsLive:           true,
			IsRecorded:       false,
			CommentCount:     0,
			LargeThumbnail:   "https://example.com/large_thumbnail.png",
			SmallThumbnail:   "https://example.com/small_thumbnail.png",
			Country:          "JP",
			Duration:         0,
			Created:          0,
			IsCollabo:        false,
			IsProtected:      false,
			MaxViewCount:     0,
			CurrentViewCount: 0,
			TotalViewCount:   0,
			HlsUrl:           "https://example.com/hls_url",
		},
		Broadcaster: twitcasting.Broadcaster{
			Id:              "123456",
			ScreenId:        "test_screen_id",
			Name:            "test_name",
			Image:           "https://example.com/image.png",
			Profile:         "test_profile",
			Level:           1,
			LastMovieId:     "12345432",
			IsLive:          true,
			SupporterCount:  0,
			SupportingCount: 0,
			Created:         0,
		},
		Tags: []string{"tag1", "tag2"},
	}
	server := CreateTestSever(t, expected, "", url.Values{}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	movieResponse, errorResponse, err := locator.Movie.GetMovie("12345432", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, movieResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	movieResponse, errorResponse, err = locator.Movie.GetMovie("12345432", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, movieResponse)
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
	movieResponse, errorResponse, err = locator.Movie.GetMovie("12345432", true)
	assert.Nil(t, movieResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestGetUserMovies(t *testing.T) {
	expected := twitcasting.UserMoviesContainer{
		Movies: []twitcasting.Movie{
			{
				Id:     "12345432",
				UserId: "123456",
			},
		},
		TotalCount: 100,
	}
	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err := locator.Movie.GetUserMovies("12345432", 10, 0, false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, moviesResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err = locator.Movie.GetUserMovies("12345432", 10, 0, true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, moviesResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err = locator.Movie.GetUserMovies("12345432", 10, 0, true)
	assert.Nil(t, moviesResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestGetUserMoviesBySliceId(t *testing.T) {
	expected := twitcasting.UserMoviesContainer{
		Movies: []twitcasting.Movie{
			{
				Id:               "123454",
				UserId:           "123456",
				Title:            "test_title",
				Category:         "test_category",
				Link:             "https://example.com/link",
				IsLive:           true,
				IsRecorded:       false,
				CommentCount:     0,
				LargeThumbnail:   "https://example.com/large_thumbnail.png",
				SmallThumbnail:   "https://example.com/small_thumbnail.png",
				Country:          "JP",
				Duration:         0,
				Created:          0,
				IsCollabo:        false,
				IsProtected:      false,
				MaxViewCount:     0,
				CurrentViewCount: 0,
				TotalViewCount:   0,
				HlsUrl:           "https://example.com/hls_url",
			},
		},
		TotalCount: 100,
	}

	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "slice_id": []string{"12345432"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err := locator.Movie.GetUserMoviesBySliceId("123454", 10, "12345432", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, moviesResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "slice_id": []string{"12345432"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err = locator.Movie.GetUserMoviesBySliceId("123454", 10, "12345432", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, moviesResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "", url.Values{"limit": []string{"10"}, "slice_id": []string{"12345432"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	moviesResponse, errorResponse, err = locator.Movie.GetUserMoviesBySliceId("123454", 10, "12345432", true)
	assert.Nil(t, moviesResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestGetCurrentLive(t *testing.T) {
	expected := twitcasting.MovieContainer{
		Movie: twitcasting.Movie{
			Id:     "12345432",
			UserId: "123456",
		},
		Broadcaster: twitcasting.Broadcaster{
			Id:              "123456",
			ScreenId:        "test_screen_id",
			Name:            "test_name",
			Image:           "https://example.com/image.png",
			Profile:         "test_profile",
			Level:           1,
			LastMovieId:     "12345432",
			IsLive:          true,
			SupporterCount:  0,
			SupportingCount: 0,
			Created:         0,
		},
		Tags: []string{"tag1", "tag2"},
	}
	server := CreateTestSever(t, expected, "", url.Values{}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	movieResponse, errorResponse, err := locator.Movie.GetCurrentLive("12345432", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, movieResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	movieResponse, errorResponse, err = locator.Movie.GetCurrentLive("12345432", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, movieResponse)
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
	movieResponse, errorResponse, err = locator.Movie.GetCurrentLive("12345432", true)
	assert.Nil(t, movieResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestPostCurrentLiveSubtitle(t *testing.T) {
	expected := twitcasting.CurrentLiveSubtitleContainer{
		MovieId:  "12345432",
		Subtitle: "test_subtitle",
	}
	server := CreateTestSever(t, expected, "{\"subtitle\":\"test_subtitle\"}", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	subtitleResponse, errorResponse, err := locator.Movie.PostCurrentLiveSubtitle("test_subtitle")
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, subtitleResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "{\"subtitle\":\"test_subtitle\"}", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	subtitleResponse, errorResponse, err = locator.Movie.PostCurrentLiveSubtitle("test_subtitle")
	assert.Nil(t, subtitleResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestDeleteCurrentLiveSubtitle(t *testing.T) {
	expected := twitcasting.CurrentLiveSubtitleContainer{
		MovieId:  "12345432",
		Subtitle: "",
	}
	server := CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	subtitleResponse, errorResponse, err := locator.Movie.DeleteCurrentLiveSubtitle()
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, subtitleResponse)
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
	subtitleResponse, errorResponse, err = locator.Movie.DeleteCurrentLiveSubtitle()
	assert.Nil(t, subtitleResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestPostCurrentLiveHashtag(t *testing.T) {
	expected := twitcasting.CurrentLiveHashtagContainer{
		MovieId: "12345432",
		Hashtag: "test_hashtag",
	}
	server := CreateTestSever(t, expected, "{\"hashtag\":\"test_hashtag\"}", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	hashtagResponse, errorResponse, err := locator.Movie.PostCurrentLiveHashtag("test_hashtag")
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, hashtagResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "{\"hashtag\":\"test_hashtag\"}", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	hashtagResponse, errorResponse, err = locator.Movie.PostCurrentLiveHashtag("test_hashtag")
	assert.Nil(t, hashtagResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestDeleteCurrentLiveHashtag(t *testing.T) {
	expected := twitcasting.CurrentLiveHashtagContainer{
		MovieId: "12345432",
		Hashtag: "",
	}
	server := CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	hashtagResponse, errorResponse, err := locator.Movie.DeleteCurrentLiveHashtag()
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, hashtagResponse)
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
	hashtagResponse, errorResponse, err = locator.Movie.DeleteCurrentLiveHashtag()
	assert.Nil(t, hashtagResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
