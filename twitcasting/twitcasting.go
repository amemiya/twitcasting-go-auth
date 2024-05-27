package twitcasting

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

const baseUrl = "https://apiv2.twitcasting.tv"

type ServiceBase struct {
	Client *Client
	Logger *Logger
}

type BasicAndBearerToken struct {
	basic  string
	bearer string
}

func (b *BasicAndBearerToken) SetBasic(basic string) {
	b.basic = basic
}

func (b *BasicAndBearerToken) SetBearer(bearer string) {
	b.bearer = bearer
}

type Client struct {
	client              *http.Client
	baseURL             string
	basicAndBearerToken BasicAndBearerToken
}

func (c *Client) SetClient(client *http.Client) {
	c.client = client
}

func (c *Client) SetBaseUrl(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetBasicAndBearerToken(basicAndBearerToken BasicAndBearerToken) {
	c.basicAndBearerToken = basicAndBearerToken
}

func (c *Client) get(path string, useBearerToken bool) (*http.Response, error) {
	request, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Version", "2.0")
	request.Header.Set("Accept", "application/json")
	if useBearerToken {
		request.Header.Set("Authorization", "Bearer "+c.basicAndBearerToken.bearer)
	} else {
		request.Header.Set("Authorization", "Basic "+c.basicAndBearerToken.basic)
	}
	response, err := c.client.Do(request)
	return response, err
}

func (c *Client) post(path string, requestBody interface{}, useBearerToken bool) (*http.Response, error) {
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", c.baseURL+path, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Version", "2.0")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	if useBearerToken {
		request.Header.Set("Authorization", "Bearer "+c.basicAndBearerToken.bearer)
	} else {
		request.Header.Set("Authorization", "Basic "+c.basicAndBearerToken.basic)
	}
	response, err := c.client.Do(request)
	return response, err
}

func (c *Client) put(path string, requestBody interface{}, useBearerToken bool) (*http.Response, error) {
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", c.baseURL+path, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Version", "2.0")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	if useBearerToken {
		request.Header.Set("Authorization", "Bearer "+c.basicAndBearerToken.bearer)
	} else {
		request.Header.Set("Authorization", "Basic "+c.basicAndBearerToken.basic)
	}
	response, err := c.client.Do(request)
	return response, err
}

func (c *Client) delete(path string, useBearerToken bool) (*http.Response, error) {
	request, err := http.NewRequest("DELETE", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Version", "2.0")
	request.Header.Set("Accept", "application/json")
	if useBearerToken {
		request.Header.Set("Authorization", "Bearer "+c.basicAndBearerToken.bearer)
	} else {
		request.Header.Set("Authorization", "Basic "+c.basicAndBearerToken.basic)
	}
	response, err := c.client.Do(request)
	return response, err
}

func (c *Client) BodyClose(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		panic(err)
	}
}

type ServiceLocator struct {
	Auth        *AuthService
	Broadcaster *BroadcastingService
	Category    *CategoryService
	Comment     *CommentService
	Gift        *GiftService
	Movie       *MovieService
	Search      *SearchService
	Supporter   *SupporterService
	User        *UserService
	Webhook     *WebhookService
}

func CreateServiceLocator(httpClient *http.Client, logger Logger, accessToken AccessToken) (*ServiceLocator, error) {
	basicAndBearerToken := BasicAndBearerToken{
		bearer: accessToken.Bearer,
		basic:  base64.StdEncoding.EncodeToString([]byte(accessToken.ClientId + ":" + accessToken.ClientSecret)),
	}
	client := &Client{
		client:              httpClient,
		baseURL:             baseUrl,
		basicAndBearerToken: basicAndBearerToken,
	}
	serviceLocator := &ServiceLocator{
		Auth:        &AuthService{Client: client, Logger: &logger},
		Broadcaster: &BroadcastingService{Client: client, Logger: &logger},
		Category:    &CategoryService{Client: client, Logger: &logger},
		Comment:     &CommentService{Client: client, Logger: &logger},
		Gift:        &GiftService{Client: client, Logger: &logger},
		Movie:       &MovieService{Client: client, Logger: &logger},
		Search:      &SearchService{Client: client, Logger: &logger},
		Supporter:   &SupporterService{Client: client, Logger: &logger},
		User:        &UserService{Client: client, Logger: &logger},
		Webhook:     &WebhookService{Client: client, Logger: &logger},
	}
	return serviceLocator, nil
}
