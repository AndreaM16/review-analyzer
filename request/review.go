package request

import (
	"net/http"
	"errors"
	"fmt"
	"encoding/json"
	"bytes"
)

type ReviewQueryParameters struct {
	Item string `url:"item"`
}

type ReviewsResponse struct {
	Reviews Reviews
	Error error
}

type ReviewRequest struct {
	Item string `json:"item,omitempty"`
	Date string `json:"date,omitempty"`
	Sentiment float64 `json:"sentiment,omitempty"`
	Stars float64 `json:"stars,omitempty"`
	Content string `json:"content"`
}

type ReviewsRequest struct {
	Item string `json:"item,omitempty"`
	Reviews []ReviewRequest `json:"reviews"`
}

func GetReviewsByItem(item string) ReviewsResponse {
	q := ReviewQueryParameters{ item }
	reviewsURL := getQueryURLByEndpointAndQueryParameters("review", q)
	response, responseError := http.Get(reviewsURL); if responseError != nil {
		return ReviewsResponse{Reviews{}, responseError}
	}
	defer response.Body.Close()
	var reviews Reviews
	_ = unmarshalHttpResponseIntoInterface(response, &reviews)
	if len(reviews.Reviews) == 0 {
		return ReviewsResponse{ reviews, errors.New(fmt.Sprintf("no reviews found for item %s", item)) }
	}
	return ReviewsResponse{ reviews, responseError }
}

func PostReviewsByItem(item string, analyzedReviews []AnalyzedReview) error {
	reviewsRequest := getReviewsRequestFromItemAndAnalyzedReviews(item, analyzedReviews)
	body, bodyError := json.Marshal(reviewsRequest); if bodyError != nil {
		fmt.Println(fmt.Sprintf("Unable to marshal reviews for item %s, got error: %s", item, bodyError.Error()))
		return bodyError
	}
	response, requestErr := http.Post(getReviewRequestQueryURL(), "application/json", bytes.NewBuffer(body)); if requestErr != nil {
		fmt.Println(fmt.Sprintf("Unable to post new reviews entries for item %s, got error: %s", item, requestErr.Error()))
		return requestErr
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println(fmt.Sprintf("Unable to post reviews entries for item %s, got status code: %d", item, response.StatusCode))
		return errors.New(fmt.Sprintf("Unable to post reviews entries for item %s", item))
	}
	fmt.Println(fmt.Sprintf("Successfully posted reviews for item %s", item))
	return nil
}

func getReviewsRequestFromItemAndAnalyzedReviews(item string, analyzedReviews []AnalyzedReview) ReviewsRequest {
	var reviewsRequest ReviewsRequest
	reviewsRequest.Item = item
	for _, r := range analyzedReviews {
		reviewsRequest.Reviews = append(reviewsRequest.Reviews, ReviewRequest{
			Content: r.Content,
			Date: r.Date,
			Item: item,
			Sentiment: r.Sentiment,
			Stars: r.Stars,
		})
	}
	return reviewsRequest
}