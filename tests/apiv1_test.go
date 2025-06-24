package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/FriendlyJJC/api_server/apiv1"
)

func TestAddUrl(t *testing.T) {
	body := apiv1.AddURLBody{LongURL: "youtube.com"}
	json_body, _ := json.Marshal(body)
	req, req_err := http.Post("http://localhost:8080/v1/shorturl/add", "application/json", strings.NewReader(string(json_body)))
	if req_err != nil {
		t.Fatalf("Something is happend with the Request: %v", req_err)
	}
	if req.StatusCode != 200 {
		t.Errorf("The Reponse Code is not 200 except it is %d", req.StatusCode)
	}
}
