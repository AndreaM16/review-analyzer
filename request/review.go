package request

import (
	"net/http"
	"errors"
	"fmt"
)

type ReviewQueryParameters struct {
	Item string `url:"item"`
}

type ReviewsResponse struct {
	Reviews Reviews `json:"reviews"`
	Error error `json:"error"`
}

func GetReviewsByItem(item string) ReviewsResponse {
	q := ReviewQueryParameters{ item }
	reviewsURL := getQueryURLByEndpointAndQueryParameters("review", q)
	response, responseError := http.Get(reviewsURL); if responseError != nil {
		return ReviewsResponse{Reviews{}, responseError}
	}
	defer response.Body.Close()
	reviews, ok := unmarshalHttpResponseIntoInterface(response).(Reviews); if !ok {
		return ReviewsResponse{Reviews{}, errors.New("unable to unmarshal reviews")}
	}
	if len(reviews.Reviews) == 0 {
		return ReviewsResponse{ reviews, errors.New(fmt.Sprintf("no reviews found for item %s", item)) }
	}
	return ReviewsResponse{ reviews, responseError }
}
