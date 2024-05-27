package twitcasting

import (
	"encoding/json"
	"errors"
)

type Gift struct {
	Id             string `json:"id"` // Since it's not unique, there's a possibility of duplication.
	Message        string `json:"message"`
	ItemImage      string `json:"item_image"`
	ItemSubImage   string `json:"item_sub_image"`
	ItemId         string `json:"item_id"`
	ItemMp         string `json:"item_mp"`
	ItemName       string `json:"item_name"`
	UserImage      string `json:"user_image"`
	UserScreenId   string `json:"user_screen_id"`
	UserScreenName string `json:"user_screen_name"`
	UserName       string `json:"user_name"`
}

type GiftContainer struct {
	SliceId string `json:"slice_id"`
	Gifts   []Gift `json:"gifts"`
}

type GiftService ServiceBase

// GetGifts https://apiv2-doc.twitcasting.tv/#get-gifts
// GetGifts Requests can only be made using a Bearer Token.
func (giftService *GiftService) GetGifts() (*GiftContainer, *ErrorResponse, error) {
	logger := *giftService.Logger
	response, err := giftService.Client.get("/gifts", true)
	if err != nil {
		logger.Error("request failed for GetGifts", err)
		return nil, nil, err
	}
	defer giftService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(GiftContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetGifts", err)
			return nil, nil, err
		}
		logger.Debug("response for GetGifts", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetGifts", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetGifts", req)
		return nil, req, errors.New("error response")
	}
}
