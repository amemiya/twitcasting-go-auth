package twitcasting

import (
	"encoding/json"
	"errors"
)

type BroadcastingUrlContainer struct {
	Enabled   bool   `json:"enabled"`
	Url       string `json:"url"`
	StreamKey string `json:"stream_key"`
}

type BroadcastingService ServiceBase

// GetRtmpUrl https://apiv2-doc.twitcasting.tv/#get-rtmp-url
// GetRtmpUrl Requests can only be made using a Bearer Token.
func (broadcastingService *BroadcastingService) GetRtmpUrl() (*BroadcastingUrlContainer, *ErrorResponse, error) {
	logger := *broadcastingService.Logger
	response, err := broadcastingService.Client.get("/rtmp_url", true)
	if err != nil {
		logger.Error("request failed for GetRtmpUrl", err)
		return nil, nil, err
	}
	defer broadcastingService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(BroadcastingUrlContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetRtmpUrl", err)
			return nil, nil, err
		}
		logger.Debug("response for GetRtmpUrl", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetRtmpUrl", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetRtmpUrl", req)
		return nil, req, errors.New("error response")
	}
}

// GetWebMUrl https://apiv2-doc.twitcasting.tv/#get-webm-url
// GetWebMUrl Requests can only be made using a Bearer Token.
func (broadcastingService *BroadcastingService) GetWebMUrl() (*BroadcastingUrlContainer, *ErrorResponse, error) {
	logger := *broadcastingService.Logger
	response, err := broadcastingService.Client.get("/webm_url", true)
	if err != nil {
		logger.Error("request failed for GetWebMUrl", err)
		return nil, nil, err
	}
	defer broadcastingService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(BroadcastingUrlContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetWebMUrl", err)
			return nil, nil, err
		}
		logger.Debug("response for GetWebMUrl", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetWebMUrl", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetWebMUrl", req)
		return nil, req, errors.New("error response")
	}
}
