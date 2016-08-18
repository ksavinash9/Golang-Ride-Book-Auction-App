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

	usersUrl = fmt.Sprintf("%s/ride", server.URL)
}

func TestCreateRide(t *testing.T) {
	rideJson := `{
	    "UserID" : 101,
	    "DriverID": 1001,
	    "RidePrice": "1024"
	}`

	reader = strings.NewReader(rideJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUpdateRide(t *testing.T) {
	rideJson := `{
	    "XYZ" : 1001,
	    "YZA": 1001,
	}`

	reader = strings.NewReader(rideJson)

	request, err := http.NewRequest("PUT", usersUrl+"?id=14", reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 400 {
		t.Error("Bad request expected: %d", res.StatusCode)
	}
}

func TestListRides(t *testing.T) {
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