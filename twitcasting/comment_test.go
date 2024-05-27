package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetComments(t *testing.T) {
	expected := twitcasting.CommentListContainer{
		Comments: []twitcasting.Comment{
			{
				Id:      "1",
				Message: "test_message",
				FromUser: twitcasting.User{
					Id:              "1",
					Name:            "test_user",
					ScreenId:        "test_screen_id",
					Image:           "https://example.com/image.png",
					Profile:         "test_profile",
					Level:           1,
					LatestMovieId:   "1456543",
					IsLive:          true,
					SupporterCount:  0,
					SupportingCount: 0,
					Created:         0,
				},
				Created: 234565432,
			},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentsResponse, errorResponse, err := locator.Comment.GetComments("1456543", 10, 0, false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, commentsResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "offset": []string{"0"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentsResponse, errorResponse, err = locator.Comment.GetComments("1456543", 10, 0, true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, commentsResponse)
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
	commentsResponse, errorResponse, err = locator.Comment.GetComments("1456543", 10, 0, true)
	assert.Nil(t, commentsResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestGetCommentsBySliceId(t *testing.T) {
	expected := twitcasting.CommentListContainer{
		Comments: []twitcasting.Comment{
			{
				Id:      "1",
				Message: "test_message",
				FromUser: twitcasting.User{
					Id:              "1",
					Name:            "test_user",
					ScreenId:        "test_screen_id",
					Image:           "https://example.com/image.png",
					Profile:         "test_profile",
					Level:           1,
					LatestMovieId:   "1456543",
					IsLive:          true,
					SupporterCount:  0,
					SupportingCount: 0,
					Created:         0,
				},
				Created: 234565432,
			},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "slice_id": []string{"10000"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentsResponse, errorResponse, err := locator.Comment.GetCommentsBySliceId("1456543", 10, "10000", false)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, commentsResponse)
	server.Close()

	server = CreateTestSever(t, expected, "", url.Values{"limit": []string{"10"}, "slice_id": []string{"10000"}}, true, http.StatusOK)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentsResponse, errorResponse, err = locator.Comment.GetCommentsBySliceId("1456543", 10, "10000", true)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, commentsResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "", url.Values{"limit": []string{"10"}, "slice_id": []string{"10000"}}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentsResponse, errorResponse, err = locator.Comment.GetCommentsBySliceId("1456543", 10, "10000", true)
	assert.Nil(t, commentsResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestPostComment(t *testing.T) {
	expected := twitcasting.CommentContainer{
		MovieId:  "1456543",
		AllCount: 1,
		Comment: twitcasting.Comment{
			Id:      "1",
			Message: "test_message",
			FromUser: twitcasting.User{
				Id:              "1",
				Name:            "test_user",
				ScreenId:        "test_screen_id",
				Image:           "https://example.com/image.png",
				Profile:         "test_profile",
				Level:           1,
				LatestMovieId:   "1456543",
				IsLive:          true,
				SupporterCount:  0,
				SupportingCount: 0,
				Created:         0,
			},
			Created: 234565432,
		},
	}
	server := CreateTestSever(t, expected, "{\"comment\":\"test_message\",\"sns\":\"none\"}", url.Values{}, true, http.StatusCreated)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentResponse, errorResponse, err := locator.Comment.PostComment("1456543", "test_message", "none")
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, commentResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{
		Error: twitcasting.Error{
			Code:    1000,
			Message: "Invalid token",
		},
	}
	server = CreateTestSever(t, expected2, "{\"comment\":\"test_message\",\"sns\":\"none\"}", url.Values{}, true, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	commentResponse, errorResponse, err = locator.Comment.PostComment("1456543", "test_message", "none")
	assert.Nil(t, commentResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestDeleteComment(t *testing.T) {
	expected := twitcasting.DeleteCommentContainer{
		CommentId: "2345432",
	}
	server := CreateTestSever(t, expected, "", url.Values{}, true, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	deleteCommentResponse, errorResponse, err := locator.Comment.DeleteComment("345676543", "2345432")
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, deleteCommentResponse)
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
	deleteCommentResponse, errorResponse, err = locator.Comment.DeleteComment("345676543", "2345432")
	assert.Nil(t, deleteCommentResponse)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error response")
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
