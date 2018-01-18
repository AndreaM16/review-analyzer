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
	itemURL := strings.Join([]string{mainURL, endpoint}, "/")
	v, _ := query.Values(queryParameters)
	return strings.Join([]string{itemURL, v.Encode()}, "?")
}

func unmarshalHttpResponseIntoInterface(r *http.Response) interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var i interface{}
	err = json.Unmarshal(body, &i)
	if err != nil {
		panic(err)
	}
	return i
}