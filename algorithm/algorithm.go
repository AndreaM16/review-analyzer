package algorithm

import (
	"fmt"
	"github.com/andream16/review-analyzer/request"
	"github.com/andream16/review-analyzer/configuration"
	"github.com/andream16/review-analyzer/exec"
	"time"
	"strconv"
	"strings"
)

const timeLayout = "2006-01-02"
var nextDay = [4]int{ 0, 0, 1, -1 }

func StartAlgorithm() error {
	configuration.InitConfiguration()
	itemsResponse := request.GetPaginatedItems(1, 100); if itemsResponse.Error != nil {
		return itemsResponse.Error
	}
	for _, item := range itemsResponse.Items.Items {
		reviewsResponse := request.GetReviewsByItem(item.Item); if reviewsResponse.Error != nil {
			fmt.Println(reviewsResponse.Error.Error())
		}
		if len(reviewsResponse.Reviews.Reviews) > 0 {
			flattenReviews := getFlattenReviews(reviewsResponse.Reviews); if len(flattenReviews) > 0 {
				fittedReviews := fitMissingReviews(flattenReviews)
				postError := request.PostReviewsByItem(item.Item, fittedReviews); if postError != nil {
					return postError
				}
			}
		}
	}
	return nil
}

func fitMissingReviews(reviews []request.AnalyzedReview) []request.AnalyzedReview {
	var fittedReviews []request.AnalyzedReview
	for i := 0; i < len(reviews); i++ {
		if i == 0 {
			fittedReviews = append(fittedReviews, reviews[i])
		}
		if i < len(reviews) - 1 {
			fittedReviews = append(fittedReviews, fitMissingDays(reviews[i], reviews[i + 1], []request.AnalyzedReview{})...)
		}
	}
	return fittedReviews
}

func fitMissingDays(lastEntry request.AnalyzedReview, finalEntry request.AnalyzedReview, reviews []request.AnalyzedReview) []request.AnalyzedReview {
	if lastEntry.Date == finalEntry.Date {
		return reviews
	}
	currentDayDateAsTime, timeParseErr := time.Parse(timeLayout, lastEntry.Date); if timeParseErr != nil {
		fmt.Println(timeParseErr.Error())
	}
	nextDayAsString := dateToString(currentDayDateAsTime.AddDate(nextDay[0], nextDay[1], nextDay[2]))
	starsAvg := (lastEntry.Stars + finalEntry.Stars) / 2
	sentimentAvg := (lastEntry.Sentiment + finalEntry.Sentiment) / 2
	currentEntry := request.AnalyzedReview{
		Date: nextDayAsString,
		Sentiment: sentimentAvg,
		Stars: starsAvg,
		Content: lastEntry.Content,
	}
	reviews = append(reviews, currentEntry)
	return fitMissingDays(currentEntry, finalEntry, reviews)
}


func getFlattenReviews(reviews request.Reviews) []request.AnalyzedReview {
	analyzedReviews := getAnalyzedReviewsFromReviews(reviews); if len(analyzedReviews) > 0 {
		return getAveragedAnalyzedReviews(analyzedReviews)
	}
	return []request.AnalyzedReview{}
}

func getAnalyzedReviewsFromReviews(reviews request.Reviews) []request.AnalyzedReview {
	var finalAnalyzedReviews []request.AnalyzedReview
	for _, r := range reviews.Reviews {
		if len(r.Content) > 0 {
			c, e := exec.GetSentimentAnalysisFromSentence(r.Content); if e == nil {
				finalAnalyzedReviews = append(finalAnalyzedReviews, request.AnalyzedReview{
					Content: r.Content,
					Date: r.Date,
					Sentiment: c,
					Stars: float64(r.Stars),
				})
			}
		}
	}
	return finalAnalyzedReviews
}

func getAveragedAnalyzedReviews(reviews []request.AnalyzedReview) []request.AnalyzedReview {
	var finalAveragedReviews []request.AnalyzedReview
	for _, r := range reviews {
		if len(finalAveragedReviews) == 0 {
			finalAveragedReviews = append(finalAveragedReviews, filterReviewsByDate(r.Date, reviews))
		} else {
			isDateAlreadyInBucket := false
			for _, rw := range finalAveragedReviews {
				if r.Date == rw.Date {
					isDateAlreadyInBucket = true
				}
			}
			if !isDateAlreadyInBucket {
				finalAveragedReviews = append(finalAveragedReviews, filterReviewsByDate(r.Date, reviews))
			}
		}
	}
	return finalAveragedReviews
}

func getAnalyzedReviewFromSliceOfAnalyzedReviews(reviews []request.AnalyzedReview) request.AnalyzedReview {
	var sentimentSum float64
	var starsSum float64
	length := float64(len(reviews))
	for _, r := range reviews {
		sentimentSum += r.Sentiment
		starsSum += r.Stars
	}
	return request.AnalyzedReview {
		Content: reviews[0].Content,
		Date: reviews[0].Date,
		Sentiment: sentimentSum / length,
		Stars: starsSum / length,
	}
}

func filterReviewsByDate(date string, reviews []request.AnalyzedReview) request.AnalyzedReview {
	var filteredReviews []request.AnalyzedReview
	for _, r := range reviews {
		if date == r.Date {
			filteredReviews = append(filteredReviews, r)
		}
	}
	if len(filteredReviews) == 1 {
		return filteredReviews[0]
	}
	return getAnalyzedReviewFromSliceOfAnalyzedReviews(filteredReviews)
}

// Returns a date in string format yy-mm-dd
func dateToString(date time.Time) string {
	var dateEntries  []string
	dateEntries = append(dateEntries, strconv.Itoa(date.Year()))
	dateEntries = append(dateEntries, strconv.Itoa(int(date.Month())))
	if len(dateEntries[1]) == 1 {
		dateEntries[1] = "0" + dateEntries[1]
	}
	dateEntries = append(dateEntries, strconv.Itoa(date.Day()))
	if len(dateEntries[2]) == 1 {
		dateEntries[2] = "0" + dateEntries[2]
	}
	return strings.Join(dateEntries, "-")
}
