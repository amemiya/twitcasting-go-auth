package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SearchUsersContainer struct {
	Users []User `json:"users"`
}

type SearchLiveMoviesContainer struct {
	Movies []MovieContainer `json:"movies"`
}

type SearchService ServiceBase

// SearchUsers https://apiv2-doc.twitcasting.tv/#search-users
func (searchService *SearchService) SearchUsers(words string, limit int, useBearerToken bool) (*SearchUsersContainer, *ErrorResponse, error) {
	logger := *searchService.Logger
	response, err := searchService.Client.get(fmt.Sprintf("/search/users?words=%v&limit=%v&lang=ja", words, limit), useBearerToken)
	if err != nil {
		logger.Error("request failed for SearchUsers", err)
		return nil, nil, err
	}
	defer searchService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(SearchUsersContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for SearchUsers", err)
			return nil, nil, err
		}
		logger.Debug("response for SearchUsers", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for SearchUsers", err)
			return nil, nil, err
		}
		logger.Debug("error response for SearchUsers", req)
		return nil, req, errors.New("error response")
	}
}

// SearchLiveMovies https://apiv2-doc.twitcasting.tv/#search-live-movies
func (searchService *SearchService) SearchLiveMovies(contextType string, context string, limit int, useBearerToken bool) (*SearchLiveMoviesContainer, *ErrorResponse, error) {
	logger := *searchService.Logger
	response, err := searchService.Client.get(fmt.Sprintf("/search/lives?type=%v&context=%v&limit=%v&lang=ja", contextType, context, limit), useBearerToken)
	if err != nil {
		logger.Error("request failed for SearchLiveMovies", err)
		return nil, nil, err
	}
	defer searchService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(SearchLiveMoviesContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for SearchLiveMovies", err)
			return nil, nil, err
		}
		logger.Debug("response for SearchLiveMovies", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for SearchLiveMovies", err)
			return nil, nil, err
		}
		logger.Debug("error response for SearchLiveMovies", req)
		return nil, req, errors.New("error response")
	}
}
