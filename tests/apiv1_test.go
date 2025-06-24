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
	//Create a Test Route that makes an Http Req to the Post Route with Value and checks if StatusCode is 200
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

func TestGetUrl(t *testing.T) {
	//Create a Test Route that makes an Http Req the GET Route and check if StatusCode is 200
	req, req_err := http.Get(TEST_URL + "/shorturl/get")
	if req_err != nil {
		t.Fatalf("Something went wrong with the Request: %v", req_err)
	}
	if req.StatusCode != 200 {
		t.Errorf("The Response Code is not 200 as expected. actual: %d", req.StatusCode)
	}
}
