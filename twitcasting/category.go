package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SubCategory struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Category struct {
	Id            string        `json:"id"`
	Name          string        `json:"name"`
	SubCategories []SubCategory `json:"sub_categories"`
}

type CategoriesContainer struct {
	Categories []Category `json:"categories"`
}

type CategoryService ServiceBase

// GetCategories https://apiv2-doc.twitcasting.tv/#get-categories
func (categoryService *CategoryService) GetCategories(lang string, useBearerToken bool) (*CategoriesContainer, *ErrorResponse, error) {
	logger := *categoryService.Logger
	response, err := categoryService.Client.get(fmt.Sprintf("/categories?lang=%v", lang), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetCategories", err)
		return nil, nil, err
	}
	defer categoryService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(CategoriesContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetCategories", err)
			return nil, nil, err
		}
		logger.Debug("response for GetCategories", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetCategories", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetCategories", req)
		return nil, req, errors.New("error response")
	}
}
