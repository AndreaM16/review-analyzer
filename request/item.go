package request

import (
	"net/http"
	"errors"
	"fmt"
)

type ItemQueryParameters struct {
	Page int `url:"page"`
	Size int `url:"size"`
}

type ItemsResponse struct {
	Items Items
	Error error
}

func GetPaginatedItems(page, size int) ItemsResponse {
	q := ItemQueryParameters{ page, size }
	mainURL := getQueryURLByEndpointAndQueryParameters("item", q)
	response, responseError := http.Get(mainURL); if responseError != nil {
		return ItemsResponse{Items{}, responseError}
	}
	defer response.Body.Close()
	var items Items
	_ = unmarshalHttpResponseIntoInterface(response, &items)
	if len(items.Items) == 0 {
		return ItemsResponse{Items{}, errors.New(fmt.Sprintf("no items found for page: %d, size: %d", page, size))}
	}
	return ItemsResponse{ items, responseError }
}
