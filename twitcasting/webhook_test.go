package twitcasting_test

import (
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetWebhookList(t *testing.T) {
	expected := twitcasting.WebhookListContainer{
		AllCount: 2,
		Webhooks: []twitcasting.Webhook{
			{UserId: "7134775954", Event: "livestart"},
			{UserId: "7134775954", Event: "liveend"},
		},
	}
	server := CreateTestSever(t, expected, "", url.Values{"limit": []string{"100"}, "offset": []string{"0"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	webhookResponse, errorResponse, err := locator.Webhook.GetWebhookList(100, 0)
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, webhookResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"limit": []string{"100"}, "offset": []string{"0"}}, false, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	webhookResponse, errorResponse, err = locator.Webhook.GetWebhookList(100, 0)
	assert.Nil(t, webhookResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestPostWebhook(t *testing.T) {
	expected := twitcasting.PostWebhookContainer{
		UserId:      "test",
		AddedEvents: []string{"livestart", "liveend"},
	}
	server := CreateTestSever(t, expected, "{\"user_id\":\"test\",\"events\":[\"livestart\",\"liveend\"]}", url.Values{}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	webhookResponse, errorResponse, err := locator.Webhook.PostWebhook("test", []string{"livestart", "liveend"})
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, webhookResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "{\"user_id\":\"test\",\"events\":[\"livestart\",\"liveend\"]}", url.Values{}, false, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	webhookResponse, errorResponse, err = locator.Webhook.PostWebhook("test", []string{"livestart", "liveend"})
	assert.Nil(t, webhookResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}

func TestDeleteWebhook(t *testing.T) {
	expected := twitcasting.DeleteWebhookContainer{
		UserId:        "test",
		DeletedEvents: []string{"livestart", "liveend"},
	}
	server := CreateTestSever(t, expected, "", url.Values{"events[]": []string{"livestart", "liveend"}, "user_id": []string{"test"}}, false, http.StatusOK)
	server.Start()
	locator := CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	webhookResponse, errorResponse, err := locator.Webhook.DeleteWebhook("test", []string{"livestart", "liveend"})
	assert.Nil(t, err)
	assert.Nil(t, errorResponse)
	assert.Equal(t, &expected, webhookResponse)
	server.Close()

	expected2 := twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "Invalid token"}}
	server = CreateTestSever(t, expected2, "", url.Values{"events[]": []string{"livestart", "liveend"}, "user_id": []string{"test"}}, false, http.StatusBadRequest)
	server.Start()
	locator = CreateTestServiceLocator(&http.Client{}, server.URL, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	webhookResponse, errorResponse, err = locator.Webhook.DeleteWebhook("test", []string{"livestart", "liveend"})
	assert.Nil(t, webhookResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "error response", err.Error())
	assert.Equal(t, &expected2, errorResponse)
	server.Close()
}
