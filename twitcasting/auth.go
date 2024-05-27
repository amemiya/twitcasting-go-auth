package twitcasting

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AccessTokenContainer struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type AuthService ServiceBase

// GetAuthorizeUrl https://apiv2-doc.twitcasting.tv/#get-authorize-url
func (authService *AuthService) GetAuthorizeUrl(clientId string, state string) string {
	return fmt.Sprintf(authService.Client.baseURL+"/oauth2/authorize?client_id=%v&response_type=code&state=%v", clientId, state)
}

// PostAccessToken https://apiv2-doc.twitcasting.tv/#get-access-token
func (authService *AuthService) PostAccessToken(clientId string, clientSecret string, code string, redirectUrl string) (*AccessTokenContainer, *ErrorResponse, error) {
	logger := *authService.Logger
	values := url.Values{
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"redirect_uri":  {redirectUrl},
	}
	request, err := http.NewRequest("POST", authService.Client.baseURL+"/oauth2/access_token", strings.NewReader(values.Encode()))
	if err != nil {
		logger.Error("create request object failed for PostAccessToken", err)
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := authService.Client.client.Do(request)
	if err != nil {
		logger.Error("request failed for PostAccessToken", err)
		return nil, nil, err
	}
	defer authService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		req := new(AccessTokenContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for PostAccessToken", err)
			return nil, nil, err
		}
		logger.Debug("response for PostAccessToken", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for PostAccessToken", err)
			return nil, nil, err
		}
		logger.Debug("error response for PostAccessToken", req)
		return nil, req, nil
	}
}
