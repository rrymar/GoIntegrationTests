package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testapi/albums"
	"testing"
)

func initTestApi() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return albums.SetupRoutes()
}

func Test_ItCreatesAlbum(t *testing.T) {
	router := initTestApi()
	server := httptest.NewServer(router)

	expected := albums.Album{
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}

	var actual albums.Album
	client := resty.New()
	resp, _ := client.R().
		SetHeader("Accept", "application/json").
		SetBody(expected).
		SetResult(&actual).
		Post(server.URL + "/albums")

	assert.Equal(t, http.StatusCreated, resp.StatusCode())

	assert.NotEmpty(t, actual.ID)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Artist, actual.Artist)
	assert.Equal(t, expected.Price, actual.Price)
}

func Test_ItReturnsAlbums(t *testing.T) {
	router := initTestApi()
	server := httptest.NewServer(router)

	var actual []albums.Album
	client := resty.New()
	resp, _ := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get(server.URL + "/albums")

	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.GreaterOrEqual(t, len(actual), 3)
}
