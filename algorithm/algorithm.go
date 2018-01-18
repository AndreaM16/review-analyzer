package algorithm

import (
	"github.com/andream16/review-analyzer/request"
	"fmt"
)

func StartAlgorithm() error {
	itemsResponse := request.GetPaginatedItems(1, 100); if itemsResponse.Error != nil {
		return itemsResponse.Error
	}
	for _, item := range itemsResponse.Items.Items {
		reviewsResponse := request.GetReviewsByItem(item.Item); if reviewsResponse.Error != nil {
			fmt.Println(reviewsResponse.Error.Error())
		}
		for idx, review := range reviewsResponse.Reviews.Reviews {
			fmt.Println(fmt.Sprintf("idx: %d, item: %s, review: %s", idx, item, review.Content))
		}
	}
}
