package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/FriendlyJJC/api_server/apiv1"
)

const TEST_URL = "http://localhost:8080/v1"

func TestAddUrl(t *testing.T) {
	//Create a Test that makes an Http Req to the Post Route with Value and checks if StatusCode is 200
	body := apiv1.AddURLBody{LongURL: "youtube.com"}
	json_body, _ := json.Marshal(body)
	req, req_err := http.Post(TEST_URL+"/shorturl/add", "application/json", strings.NewReader(string(json_body)))
	if req_err != nil {
		t.Fatalf("Something went wrong with the Request: %v", req_err)
	}
	if req.StatusCode != 200 {
		t.Errorf("The Reponse Code is not 200 as expected. actual: %d and url was %s", req.StatusCode, req.Request.URL)
	}
}

func TestGetUrls(t *testing.T) {
	//Create a Test that makes an Http Req the GET Route and check if StatusCode is 200
	req, req_err := http.Get(TEST_URL + "/shorturl/get")
	if req_err != nil {
		t.Fatalf("Something went wrong with the Request: %v", req_err)
	}
	if req.StatusCode != 200 {
		t.Errorf("The Response Code is not 200 as expected. actual: %d", req.StatusCode)
	}
}

func TestGetUrl(t *testing.T) {
	var TEST_SHORTURL string
	//Create a Test that makes an HTTP Req to the GET/:ID Route to Get a Specific URL from the shorturl
	req, req_error := http.Get(TEST_URL + "/shorturl/get")
	if req_error != nil {
		t.Fatalf("Something went wrong with the GET Req to get all IDS: %v", req_error)
	}
	var ResponseBody apiv1.ShortURLS
	json.NewDecoder(req.Body).Decode(&ResponseBody)
	TEST_SHORTURL = *ResponseBody.Data[0].ShortURL
	req_2, req_2_error := http.Get(TEST_URL + "/shorturl/get/" + TEST_SHORTURL)
	if req_2_error != nil {
		t.Fatalf("Something went wrong with the GET Req with the shorturl")
	}
	if req_2.StatusCode != 200 {
		t.Errorf("The Response Code is not 200 as expected. actual: %d", req_2.StatusCode)
	}
}
