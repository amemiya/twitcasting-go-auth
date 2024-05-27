package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Comment struct {
	Id       string `json:"id"`
	Message  string `json:"message"`
	FromUser User   `json:"from_user"`
	Created  int    `json:"created"`
}

type CommentListContainer struct {
	MovieId  string    `json:"movie_id"`
	AllCount int       `json:"all_count"`
	Comments []Comment `json:"comments"`
}

type CommentContainer struct {
	MovieId  string  `json:"movie_id"`
	AllCount int     `json:"all_count"`
	Comment  Comment `json:"comment"`
}

type CommentRequestBody struct {
	Comment string `json:"comment"`
	Sns     string `json:"sns"`
}

type DeleteCommentContainer struct {
	CommentId string `json:"comment_id"`
}

type CommentService ServiceBase

// GetComments https://apiv2-doc.twitcasting.tv/#get-comments
func (commentService *CommentService) GetComments(movieId string, limit int, offset int, useBearerToken bool) (*CommentListContainer, *ErrorResponse, error) {
	logger := *commentService.Logger
	response, err := commentService.Client.get(fmt.Sprintf("/movies/%v/comments?limit=%v&offset=%v", movieId, limit, offset), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetComments", err)
		return nil, nil, err
	}
	defer commentService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(CommentListContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetComments", err)
			return nil, nil, err
		}
		logger.Debug("response for GetComments", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetComments", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetComments", req)
		return nil, req, errors.New("error response")
	}
}

// GetCommentsBySliceId https://apiv2-doc.twitcasting.tv/#get-comments
func (commentService *CommentService) GetCommentsBySliceId(movieId string, limit int, sliceId string, useBearerToken bool) (*CommentListContainer, *ErrorResponse, error) {
	logger := *commentService.Logger
	response, err := commentService.Client.get(fmt.Sprintf("/movies/%v/comments?limit=%v&slice_id=%v", movieId, limit, sliceId), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetCommentsBySliceId", err)
		return nil, nil, err
	}
	defer commentService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(CommentListContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetCommentsBySliceId", err)
			return nil, nil, err
		}
		logger.Debug("response for GetCommentsBySliceId", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetCommentsBySliceId", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetCommentsBySliceId", req)
		return nil, req, errors.New("error response")
	}
}

// PostComment @see https://apiv2-doc.twitcasting.tv/#post-comment
// PostComment Requests can only be made using a Bearer Token.
func (commentService *CommentService) PostComment(movieId string, message string, sns string) (*CommentContainer, *ErrorResponse, error) {
	logger := *commentService.Logger
	response, err := commentService.Client.post(fmt.Sprintf("/movies/%v/comments", movieId), CommentRequestBody{Comment: message, Sns: sns}, true)
	if err != nil {
		logger.Error("request failed for PostComment", err)
		return nil, nil, err
	}
	defer commentService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		req := new(CommentContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for PostComment", err)
			return nil, nil, err
		}
		logger.Debug("response for PostComment", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for PostComment", err)
			return nil, nil, err
		}
		logger.Debug("error response for PostComment", req)
		return nil, req, errors.New("error response")
	}
}

// DeleteComment @see https://apiv2-doc.twitcasting.tv/#delete-comment
// DeleteComment Requests can only be made using a Bearer Token.
func (commentService *CommentService) DeleteComment(movieId string, commentId string) (*DeleteCommentContainer, *ErrorResponse, error) {
	logger := *commentService.Logger
	response, err := commentService.Client.delete(fmt.Sprintf("/movies/%v/comments/%v", movieId, commentId), true)
	if err != nil {
		logger.Error("request failed for DeleteComment", err)
		return nil, nil, err
	}
	defer commentService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(DeleteCommentContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for DeleteComment", err)
			return nil, nil, err
		}
		logger.Debug("response for DeleteComment", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for DeleteComment", err)
			return nil, nil, err
		}
		logger.Debug("error response for DeleteComment", req)
		return nil, req, errors.New("error response")
	}
}
