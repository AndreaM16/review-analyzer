package request

import (
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/google/go-querystring/query"
	"github.com/andream16/review-analyzer/configuration"
)

func getQueryURLByEndpointAndQueryParameters(endpoint string, queryParameters interface{}) string {
	config := configuration.GetConfiguration()
	mainURL := strings.Join([]string{config.REMOTE.HOST, config.REMOTE.PORT}, ":")
	itemURL := strings.Join([]string{mainURL, config.REMOTE.ENDPOINTS.BASE, endpoint}, "/")
	v, _ := query.Values(queryParameters)
	return strings.Join([]string{itemURL, v.Encode()}, "?")
}

func getReviewRequestQueryURL() string {
	config := configuration.GetConfiguration()
	mainURL := strings.Join([]string{config.REMOTE.HOST, config.REMOTE.PORT}, ":")
	return strings.Join([]string{mainURL, config.REMOTE.ENDPOINTS.BASE, config.REMOTE.ENDPOINTS.REVIEW}, "/")
}

func unmarshalHttpResponseIntoInterface(r *http.Response, v interface{}) interface{} {
	body, readErr := ioutil.ReadAll(r.Body); if readErr != nil {
		panic(readErr)
	}
	unmarshalErr := json.Unmarshal(body, v); if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	return v
}