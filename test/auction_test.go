package test

import (
	"../routers"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	usersUrl string
)

func init() {

	router := routers.GetRouter()

	server = httptest.NewServer(router)

	usersUrl = fmt.Sprintf("%s/auction", server.URL)
}

func TestCreateAuction(t *testing.T) {
	auctionJson := `{
	    "UserID" : 101,
	    "DriverID": 1001,
	    "RidePrice": "1024"
	}`

	reader = strings.NewReader(auctionJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUpdateAuction(t *testing.T) {
	auctionJson := `{
	    "XYZ" : 1001,
	    "YZA": 1001,
	}`

	reader = strings.NewReader(auctionJson)

	request, err := http.NewRequest("PUT", usersUrl+"?id=14", reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 400 {
		t.Error("Bad request expected: %d", res.StatusCode)
	}
}

func TestListAuctions(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}