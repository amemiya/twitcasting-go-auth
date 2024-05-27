package twitcasting_test

import (
	"encoding/base64"
	"encoding/json"
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

type BasicLogger struct{}

func (basicLogger *BasicLogger) Debug(v ...interface{}) {}
func (basicLogger *BasicLogger) Info(v ...interface{})  {}
func (basicLogger *BasicLogger) Warn(v ...interface{})  {}
func (basicLogger *BasicLogger) Error(v ...interface{}) {}

func CreateTestServiceLocator(httpClient *http.Client, urlParse string, accessToken twitcasting.AccessToken) *twitcasting.ServiceLocator {
	basicAndBearerToken := twitcasting.BasicAndBearerToken{}
	basicAndBearerToken.SetBearer(accessToken.Bearer)
	basicAndBearerToken.SetBasic(base64.StdEncoding.EncodeToString([]byte(accessToken.ClientId + ":" + accessToken.ClientSecret)))
	client := twitcasting.Client{}
	client.SetClient(httpClient)
	client.SetBaseUrl(urlParse)
	client.SetBasicAndBearerToken(basicAndBearerToken)
	var logger twitcasting.Logger = &BasicLogger{}
	serviceLocator := &twitcasting.ServiceLocator{
		Auth:        &twitcasting.AuthService{Client: &client, Logger: &logger},
		Broadcaster: &twitcasting.BroadcastingService{Client: &client, Logger: &logger},
		Category:    &twitcasting.CategoryService{Client: &client, Logger: &logger},
		Comment:     &twitcasting.CommentService{Client: &client, Logger: &logger},
		Gift:        &twitcasting.GiftService{Client: &client, Logger: &logger},
		Movie:       &twitcasting.MovieService{Client: &client, Logger: &logger},
		Search:      &twitcasting.SearchService{Client: &client, Logger: &logger},
		Supporter:   &twitcasting.SupporterService{Client: &client, Logger: &logger},
		User:        &twitcasting.UserService{Client: &client, Logger: &logger},
		Webhook:     &twitcasting.WebhookService{Client: &client, Logger: &logger},
	}
	return serviceLocator
}

func CreateTestSever(t *testing.T, resBody interface{}, reqBodyParams string, reqQueryParams url.Values, useBearerToken bool, statusCode int) *httptest.Server {
	authorization := "Basic Y2xpZW50OnNlY3JldA=="
	if useBearerToken {
		authorization = "Bearer bearer"
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpHeader := http.Header{}
		if r.Method == "GET" || r.Method == "DELETE" {
			httpHeader = http.Header{
				"Accept":          {"application/json"},
				"Accept-Encoding": {"gzip"},
				"Authorization":   {authorization},
				"User-Agent":      {"Go-http-client/1.1"},
				"X-Api-Version":   {"2.0"},
			}
		} else if r.Method == "POST" || r.Method == "PUT" {
			httpHeader = http.Header{
				"Accept":          {"application/json"},
				"Accept-Encoding": {"gzip"},
				"Authorization":   {authorization},
				"Content-Length":  {strconv.FormatInt(r.ContentLength, 10)},
				"Content-Type":    {"application/json"},
				"User-Agent":      {"Go-http-client/1.1"},
				"X-Api-Version":   {"2.0"},
			}
		} else {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		assert.Equal(t, httpHeader, r.Header)

		// Request body is checked
		b, _ := io.ReadAll(r.Body)
		assert.Equal(t, reqBodyParams, string(b))

		// Request query params is checked
		assert.Equal(t, r.URL.Query(), reqQueryParams)

		body, _ := json.Marshal(resBody)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		if statusCode == 200 || statusCode == 201 {
			w.Header().Set("X-Ratelimit-Limit", "60")
			w.Header().Set("X-Ratelimit-Remaining", "59")
			w.Header().Set("X-Ratelimit-Reset", "1716600000")
		}
		w.Header().Set("X-We-Are-Hiring", "https://about.moi.st/ja/recruit/")
		w.WriteHeader(statusCode)
		writeString, _ := io.WriteString(w, string(body))
		assert.Equal(t, len(body), writeString)
	})
	listener, _ := net.Listen("tcp", ":8080")
	return &httptest.Server{Listener: listener, Config: &http.Server{Handler: handler}}
}

func TestCreateServiceLocator(t *testing.T) {
	locator, _ := twitcasting.CreateServiceLocator(nil, nil, twitcasting.AccessToken{ClientId: "client", ClientSecret: "secret", Bearer: "bearer"})
	// locatorがnilでないことを確認
	assert.NotNil(t, locator.Auth)
	assert.NotNil(t, locator.Broadcaster)
	assert.NotNil(t, locator.Category)
	assert.NotNil(t, locator.Comment)
	assert.NotNil(t, locator.Gift)
	assert.NotNil(t, locator.Movie)
	assert.NotNil(t, locator.Search)
	assert.NotNil(t, locator.Supporter)
	assert.NotNil(t, locator.User)
	assert.NotNil(t, locator.Webhook)
}

func TestHappyPath(t *testing.T) {
	t.Skip("Skip TestHappyPath")
	accessToken := twitcasting.AccessToken{}
	locator, _ := twitcasting.CreateServiceLocator(&http.Client{}, &twitcasting.BasicLogger{Logger: &log.Logger{}}, accessToken)

	rtmpUrl, errorResponse, err := locator.Broadcaster.GetRtmpUrl()
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.True(t, rtmpUrl.Enabled)
	assert.NotEqual(t, "", rtmpUrl.Url)
	assert.NotEqual(t, "", rtmpUrl.StreamKey)

	webMUrl, errorResponse, err := locator.Broadcaster.GetWebMUrl()
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.True(t, webMUrl.Enabled)
	assert.NotEqual(t, "", webMUrl.Url)
	assert.Equal(t, "", webMUrl.StreamKey)

	// test categories
	categories, errorResponse, err := locator.Category.GetCategories("ja", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(categories.Categories))
	categories2, errorResponse, err := locator.Category.GetCategories("ja", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, categories, categories2)

	// test comments
	comments, errorResponse, err := locator.Comment.GetComments("783302520", 10, 0, false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 9, len(comments.Comments))
	comments2, errorResponse, err := locator.Comment.GetComments("783302520", 10, 0, true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, comments, comments2)

	comments3, errorResponse, err := locator.Comment.GetComments("783302520", 10, 10, false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 9, len(comments3.Comments))
	comments4, errorResponse, err := locator.Comment.GetComments("783302520", 10, 10, true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, comments3, comments4)

	comments5, errorResponse, err := locator.Comment.GetCommentsBySliceId("783302520", 10, "28622656685", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(comments5.Comments))
	comments6, errorResponse, err := locator.Comment.GetCommentsBySliceId("783302520", 10, "28622656685", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, comments5, comments6)

	comments7, errorResponse5, err5 := locator.Comment.PostComment("793641303", "test message", "none")
	assert.Nil(t, errorResponse5)
	assert.Nil(t, err5)
	assert.Equal(t, "test message", comments7.Comment.Message)
	comments8, errorResponse6, err6 := locator.Comment.DeleteComment("793641303", comments7.Comment.Id)
	assert.Nil(t, errorResponse6)
	assert.Nil(t, err6)
	assert.Equal(t, comments7.Comment.Id, comments8.CommentId)

	// test gifts
	gifts, errorResponse, err := locator.Gift.GetGifts()
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(gifts.Gifts))

	// test movies
	movie, errorResponse, err := locator.Movie.GetMovie("783302520", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "783302520", movie.Movie.Id)
	assert.Equal(t, "1284086739074560001", movie.Movie.UserId)
	assert.Equal(t, "こんばんはっ！", movie.Movie.Title)
	assert.Equal(t, "生きてます", movie.Movie.Subtitle)
	assert.Equal(t, "こんばんはっ！", movie.Movie.LastOwnerComment)
	assert.Equal(t, "girls_cutevoice_jp", movie.Movie.Category)
	assert.Equal(t, "https://twitcasting.tv/ichineko__/movie/783302520", movie.Movie.Link)
	assert.False(t, movie.Movie.IsLive)
	assert.False(t, movie.Movie.IsRecorded)
	assert.Equal(t, 641, movie.Movie.CommentCount)
	assert.NotEqual(t, "", movie.Movie.LargeThumbnail)
	assert.NotEqual(t, "", movie.Movie.SmallThumbnail)
	assert.Equal(t, "jp", movie.Movie.Country)
	assert.Equal(t, 3664, movie.Movie.Duration)
	assert.Equal(t, 1702905255, movie.Movie.Created)
	assert.False(t, movie.Movie.IsCollabo)
	assert.False(t, movie.Movie.IsProtected)
	assert.Equal(t, 37, movie.Movie.MaxViewCount)
	assert.Equal(t, 0, movie.Movie.CurrentViewCount)
	assert.Equal(t, 264, movie.Movie.TotalViewCount)
	movie2, errorResponse, err := locator.Movie.GetMovie("783302520", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, movie, movie2)

	movies, errorResponse, err := locator.Movie.GetUserMovies("ichineko__", 10, 0, false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 227, movies.TotalCount)
	movies2, errorResponse, err := locator.Movie.GetUserMovies("ichineko__", 10, 0, true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, movies, movies2)

	movies3, errorResponse, err := locator.Movie.GetUserMoviesBySliceId("ichineko__", 10, "630455865", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(movies3.Movies))
	movies4, errorResponse, err := locator.Movie.GetUserMoviesBySliceId("ichineko__", 10, "630455865", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, movies3, movies4)

	// now live movie is required for this test
	currentMovie, errorResponse, err := locator.Movie.GetCurrentLive("ichineko__", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "ichineko__", currentMovie.Broadcaster.ScreenId)
	assert.Equal(t, "630455865", currentMovie.Movie.Id)
	assert.NotEqual(t, 0, currentMovie.Movie.CurrentViewCount)
	currentMovie2, errorResponse, err := locator.Movie.GetCurrentLive("ichineko__", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, currentMovie, currentMovie2)

	// now live myself movie is required for this test
	subtitle, errorResponse, err := locator.Movie.PostCurrentLiveSubtitle("test subtitle2")
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.NotEqual(t, "", subtitle.MovieId)
	assert.Equal(t, "test subtitle2", subtitle.Subtitle)
	time.Sleep(3 * time.Second) // wait for subtitle to be updated
	subtitle2, errorResponse, err := locator.Movie.DeleteCurrentLiveSubtitle()
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, subtitle.MovieId, subtitle2.MovieId)
	assert.Equal(t, "", subtitle2.Subtitle)
	time.Sleep(1 * time.Second) // wait for subtitle to be updated

	hashtag, errorResponse, err := locator.Movie.PostCurrentLiveHashtag("#初見さん大歓迎")
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.NotEqual(t, "", hashtag.MovieId)
	assert.Equal(t, "#初見さん大歓迎", hashtag.Hashtag)
	hashtag2, errorResponse, err := locator.Movie.DeleteCurrentLiveHashtag()
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, hashtag.MovieId, hashtag2.MovieId)
	assert.Equal(t, "", hashtag2.Hashtag)

	// test search
	search, errorResponse, err := locator.Search.SearchUsers("いちねこ", 10, false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(search.Users))
	assert.Equal(t, "1284086739074560001", search.Users[0].Id)
	search2, errorResponse, err := locator.Search.SearchUsers("いちねこ", 10, true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, search, search2)

	search3, errorResponse, err := locator.Search.SearchLiveMovies("recommend", "", 10, false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(search3.Movies))
	search4, errorResponse, err := locator.Search.SearchLiveMovies("recommend", "", 10, true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, search3, search4)

	// test supporter
	status, errorResponse, err := locator.Supporter.GetSupportingStatus("c:uonq", "ichineko__", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.False(t, status.IsSupporting)
	assert.Equal(t, 0, status.Supported)
	status2, errorResponse, err := locator.Supporter.GetSupportingStatus("c:uonq", "ichineko__", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, status, status2)

	support, errorResponse, err := locator.Supporter.PostSupport([]string{"c:uonq", "ichineko__"})
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 1, support.AddedCount)
	support2, errorResponse, err := locator.Supporter.DeleteSupport([]string{"c:uonq"})
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 1, support2.RemovedCount)

	supportingList, errorResponse, err := locator.Supporter.GetSupportingList("ichineko__", 10, 0, false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 128, supportingList.Total)
	supportingList2, errorResponse, err := locator.Supporter.GetSupportingList("ichineko__", 10, 0, true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, supportingList, supportingList2)

	supporterList, errorResponse, err := locator.Supporter.GetSupporterList("ichineko__", 10, 0, "new", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 2678, supporterList.Total)
	supporterList2, errorResponse, err := locator.Supporter.GetSupporterList("ichineko__", 10, 0, "new", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, supporterList, supporterList2)

	// test User
	user, errorResponse, err := locator.User.GetUser("ichineko__", false)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "1284086739074560001", user.User.Id)
	assert.NotEqual(t, "", user.User.Name)
	assert.Equal(t, "ichineko__", user.User.ScreenId)
	assert.NotEqual(t, "", user.User.Image)
	assert.Equal(t, "", user.User.Profile)
	assert.Equal(t, "", user.User.LatestMovieId)
	assert.True(t, 36 < user.User.Level)
	assert.Equal(t, "", user.User.LatestMovieId)
	assert.False(t, user.User.IsLive)
	assert.True(t, user.SupporterCount >= 0)
	assert.True(t, user.SupportingCount >= 0)

	user2, errorResponse, err := locator.User.GetUser("ichineko__", true)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, user, user2)

	webhook, errorResponse, err := locator.Webhook.PostWebhook("1284086739074560001", []string{"livestart", "liveend"})
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "1284086739074560001", webhook.UserId)
	assert.Equal(t, []string{"livestart", "liveend"}, webhook.AddedEvents)

	webhook2, errorResponse, err := locator.Webhook.GetWebhookList(10, 0)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 2, webhook2.AllCount)
	assert.Equal(t, "1284086739074560001", webhook2.Webhooks[0].UserId)
	assert.Equal(t, "1284086739074560001", webhook2.Webhooks[1].UserId)

	webhook3, errorResponse, err := locator.Webhook.DeleteWebhook("1284086739074560001", []string{"livestart", "liveend"})
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "1284086739074560001", webhook3.UserId)
	assert.Equal(t, []string{"livestart", "liveend"}, webhook3.DeletedEvents)

	webhook4, errorResponse, err := locator.Webhook.GetWebhookList(10, 0)
	assert.Nil(t, errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, 0, webhook4.AllCount)
}
