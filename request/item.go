package request

import (
	"net/http"
	"errors"
)

type ItemQueryParameters struct {
	Page int `url:"page"`
	Size int `url:"size"`
}

type ItemsResponse struct {
	Items Items `json:"items"`
	Error error `json:"error"`
}

func GetPaginatedItems(page, size int) ItemsResponse {
	q := ItemQueryParameters{ page, size }
	mainURL := getQueryURLByEndpointAndQueryParameters("item", q)
	response, responseError := http.Get(mainURL); if responseError != nil {
		return ItemsResponse{Items{}, responseError}
	}
	defer response.Body.Close()
	items, ok := unmarshalHttpResponseIntoInterface(response).(Items); if !ok || len(items.Items) == 0 {
		return ItemsResponse{Items{}, errors.New("unable to unmarshal items")}
	}
	return ItemsResponse{ items, responseError }
}
