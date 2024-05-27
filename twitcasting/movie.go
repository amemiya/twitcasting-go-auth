package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MovieContainer struct {
	Movie       Movie       `json:"movie"`
	Broadcaster Broadcaster `json:"broadcaster"`
	Tags        []string    `json:"tags"`
}

type UserMoviesContainer struct {
	Movies     []Movie `json:"movies"`
	TotalCount int     `json:"total_count"`
}

type CurrentLiveSubtitleRequestBody struct {
	Subtitle string `json:"subtitle"`
}

type CurrentLiveSubtitleContainer struct {
	MovieId  string `json:"movie_id"`
	Subtitle string `json:"subtitle"`
}

type CurrentLiveHashtagRequestBody struct {
	Hashtag string `json:"hashtag"`
}

type CurrentLiveHashtagContainer struct {
	MovieId string `json:"movie_id"`
	Hashtag string `json:"hashtag"`
}

type Movie struct {
	Id               string `json:"id"`
	UserId           string `json:"user_id"`
	Title            string `json:"title"`
	Subtitle         string `json:"subtitle"`
	LastOwnerComment string `json:"last_owner_comment"`
	Category         string `json:"category"`
	Link             string `json:"link"`
	IsLive           bool   `json:"is_live"`
	IsRecorded       bool   `json:"is_recorded"`
	CommentCount     int    `json:"comment_count"`
	LargeThumbnail   string `json:"large_thumbnail"`
	SmallThumbnail   string `json:"small_thumbnail"`
	Country          string `json:"country"`
	Duration         int    `json:"duration"`
	Created          int    `json:"created"`
	IsCollabo        bool   `json:"is_collabo"`
	IsProtected      bool   `json:"is_protected"`
	MaxViewCount     int    `json:"max_view_count"`
	CurrentViewCount int    `json:"current_view_count"`
	TotalViewCount   int    `json:"total_view_count"`
	HlsUrl           string `json:"hls_url"`
}

type Broadcaster struct {
	Id              string `json:"id"`
	ScreenId        string `json:"screen_id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Profile         string `json:"profile"`
	Level           int    `json:"level"`
	LastMovieId     string `json:"last_movie_id"`
	IsLive          bool   `json:"is_live"`
	SupporterCount  int    `json:"supporter_count"`
	SupportingCount int    `json:"supporting_count"`
	Created         int    `json:"created"`
}

type MovieService ServiceBase

// GetMovie @see https://apiv2-doc.twitcasting.tv/#movie
func (movieService *MovieService) GetMovie(movieId string, useBearerToken bool) (*MovieContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.get(fmt.Sprintf("/movies/%v", movieId), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetMovie", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(MovieContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetMovie", err)
			return nil, nil, err
		}
		logger.Debug("response for GetMovie", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetMovie", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetMovie", req)
		return nil, req, errors.New("error response")
	}
}

// GetUserMovies @see https://apiv2-doc.twitcasting.tv//#get-movies-by-user
func (movieService *MovieService) GetUserMovies(userId string, limit int, offset int, useBearerToken bool) (*UserMoviesContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.get(fmt.Sprintf("/users/%v/movies?limit=%v&offset=%v", userId, limit, offset), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetUserMovies", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(UserMoviesContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetUserMovies", err)
			return nil, nil, err
		}
		logger.Debug("response for GetUserMovies", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetUserMovies", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetUserMovies", req)
		return nil, req, errors.New("error response")
	}
}

// GetUserMoviesBySliceId @see https://apiv2-doc.twitcasting.tv//#get-movies-by-user
func (movieService *MovieService) GetUserMoviesBySliceId(userId string, limit int, sliceId string, useBearerToken bool) (*UserMoviesContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.get(fmt.Sprintf("/users/%v/movies?limit=%v&slice_id=%v", userId, limit, sliceId), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetUserMoviesBySliceId", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(UserMoviesContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetUserMoviesBySliceId", err)
			return nil, nil, err
		}
		logger.Debug("response for GetUserMoviesBySliceId", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetUserMoviesBySliceId", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetUserMoviesBySliceId", req)
		return nil, req, errors.New("error response")
	}
}

// GetCurrentLive @see https://apiv2-doc.twitcasting.tv/#get-current-live
func (movieService *MovieService) GetCurrentLive(userId string, useBearerToken bool) (*MovieContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.get(fmt.Sprintf("/users/%v/current_live", userId), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetCurrentLive", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(MovieContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetCurrentLive", err)
			return nil, nil, err
		}
		logger.Debug("response for GetCurrentLive", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetCurrentLive", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetCurrentLive", req)
		return nil, req, errors.New("error response")
	}
}

// PostCurrentLiveSubtitle @see https://apiv2-doc.twitcasting.tv/#set-current-live-subtitle
// PostCurrentLiveSubtitle Requests can only be made using a Bearer Token.
func (movieService *MovieService) PostCurrentLiveSubtitle(subtitle string) (*CurrentLiveSubtitleContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.post(
		"/movies/subtitle",
		CurrentLiveSubtitleRequestBody{Subtitle: subtitle},
		true,
	)
	if err != nil {
		logger.Error("request failed for PostCurrentLiveSubtitle", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		req := new(CurrentLiveSubtitleContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for PostCurrentLiveSubtitle", err)
			return nil, nil, err
		}
		logger.Debug("response for PostCurrentLiveSubtitle", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for PostCurrentLiveSubtitle", err)
			return nil, nil, err
		}
		logger.Debug("error response for PostCurrentLiveSubtitle", req)
		return nil, req, errors.New("error response")
	}
}

// DeleteCurrentLiveSubtitle @see https://apiv2-doc.twitcasting.tv/#unset-current-live-subtitle
// DeleteCurrentLiveSubtitle Requests can only be made using a Bearer Token.
func (movieService *MovieService) DeleteCurrentLiveSubtitle() (*CurrentLiveSubtitleContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.delete(
		"/movies/subtitle",
		true,
	)
	if err != nil {
		logger.Error("request failed for DeleteCurrentLiveSubtitle", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(CurrentLiveSubtitleContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for DeleteCurrentLiveSubtitle", err)
			return nil, nil, err
		}
		logger.Debug("response for DeleteCurrentLiveSubtitle", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for DeleteCurrentLiveSubtitle", err)
			return nil, nil, err
		}
		logger.Debug("error response for DeleteCurrentLiveSubtitle", req)
		return nil, req, errors.New("error response")
	}
}

// PostCurrentLiveHashtag @see https://apiv2-doc.twitcasting.tv/#set-current-live-hashtag
// PostCurrentLiveHashtag Requests can only be made using a Bearer Token.
func (movieService *MovieService) PostCurrentLiveHashtag(hashtag string) (*CurrentLiveHashtagContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.post(
		"/movies/hashtag",
		CurrentLiveHashtagRequestBody{Hashtag: hashtag},
		true,
	)
	if err != nil {
		logger.Error("request failed for PostCurrentLiveHashtag", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		req := new(CurrentLiveHashtagContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for PostCurrentLiveHashtag", err)
			return nil, nil, err
		}
		logger.Debug("response for PostCurrentLiveHashtag", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for PostCurrentLiveHashtag", err)
			return nil, nil, err
		}
		logger.Debug("error response for PostCurrentLiveHashtag", req)
		return nil, req, errors.New("error response")
	}
}

// DeleteCurrentLiveHashtag @see https://apiv2-doc.twitcasting.tv/#unset-current-live-hashtag
// DeleteCurrentLiveHashtag Requests can only be made using a Bearer Token.
func (movieService *MovieService) DeleteCurrentLiveHashtag() (*CurrentLiveHashtagContainer, *ErrorResponse, error) {
	logger := *movieService.Logger
	response, err := movieService.Client.delete(
		"/movies/hashtag",
		true,
	)
	if err != nil {
		logger.Error("request failed for DeleteCurrentLiveHashtag", err)
		return nil, nil, err
	}
	defer movieService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(CurrentLiveHashtagContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for DeleteCurrentLiveHashtag", err)
			return nil, nil, err
		}
		logger.Debug("response for DeleteCurrentLiveHashtag", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for DeleteCurrentLiveHashtag", err)
			return nil, nil, err
		}
		logger.Debug("error response for DeleteCurrentLiveHashtag", req)
		return nil, req, errors.New("error response")
	}
}
