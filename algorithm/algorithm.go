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

type AnalyzedReviews struct {
	Reviews []AnalyzedReview `json:"reviews"`
}

type AnalyzedReview struct {
	Date string `json:"date"`
	Sentiment float64 `json:"sentiment"`
	Stars float64 `json:"stars"`
}

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
				fmt.Println(fittedReviews)
			}
		}
	}
	return nil
}

func fitMissingReviews(reviews []AnalyzedReview) []AnalyzedReview {
	fittedReviews := make([]AnalyzedReview, 0)
	for i := 0; i <= len(reviews); i++ {
		fittedReviews = append(fittedReviews, fitMissingDays(reviews[i], reviews[i+1], []AnalyzedReview{})...)
	}
	return fittedReviews
}

func fitMissingDays(lastEntry AnalyzedReview, finalEntry AnalyzedReview, reviews []AnalyzedReview) []AnalyzedReview {
	if lastEntry.Date == finalEntry.Date {
		return reviews
	}
	currentDayDateAsTime, _ := time.Parse(timeLayout, lastEntry.Date)
	nextDay := dateToString(currentDayDateAsTime.AddDate(nextDay[0], nextDay[1], nextDay[2]))
	starsAvg := (lastEntry.Stars + finalEntry.Stars) / 2
	sentimentAvg := (lastEntry.Sentiment + finalEntry.Sentiment) / 2
	currentEntry := AnalyzedReview{
		Date: nextDay,
		Sentiment: sentimentAvg,
		Stars: starsAvg,
	}
	reviews = append(reviews, currentEntry)
	return fitMissingDays(currentEntry, finalEntry, reviews)
}


func getFlattenReviews(reviews request.Reviews) []AnalyzedReview {
	analyzedReviews := getAnalyzedReviewsFromReviews(reviews); if len(analyzedReviews) > 0 {
		return getAveragedAnalyzedReviews(analyzedReviews)
	}
	return []AnalyzedReview{}
}

func getAnalyzedReviewsFromReviews(reviews request.Reviews) []AnalyzedReview {
	var finalAnalyzedReviews []AnalyzedReview
	for _, r := range reviews.Reviews {
		if len(r.Content) > 0 {
			c, e := exec.GetSentimentAnalysisFromSentence(r.Content); if e == nil {
				finalAnalyzedReviews = append(finalAnalyzedReviews, AnalyzedReview{r.Date, c, float64(r.Stars)})
			}
		}
	}
	return finalAnalyzedReviews
}

func getAveragedAnalyzedReviews(reviews []AnalyzedReview) []AnalyzedReview {
	var finalAveragedReviews []AnalyzedReview
	for _, r := range reviews {
		f := filterReviewsByDate(r.Date, reviews); if len(f) == 1 {
			finalAveragedReviews = append(finalAveragedReviews, r)
		} else {
			finalAveragedReviews = append(finalAveragedReviews, getAnalyzedReviewFromSliceOfAnalyzedReviews(f))
		}
	}
	return finalAveragedReviews
}

func getAnalyzedReviewFromSliceOfAnalyzedReviews(reviews []AnalyzedReview) AnalyzedReview {
	var sentimentSum float64
	var starsSum float64
	length := float64(len(reviews))
	for _, r := range reviews {
		sentimentSum += r.Sentiment
		starsSum += r.Stars
	}
	return AnalyzedReview {
		Date: reviews[0].Date,
		Sentiment: sentimentSum / length,
		Stars: starsSum / length,
	}
}

func filterReviewsByDate(date string, reviews []AnalyzedReview) []AnalyzedReview {
	var filteredReviews []AnalyzedReview
	for _, r := range reviews {
		if date == r.Date {
			filteredReviews = append(filteredReviews, r)
		}
	}
	return filteredReviews
}

// Returns a date in string format yy-mm-dd
func dateToString(date time.Time) string {
	dateEntries := make([]string, 3)
	dateEntries[0] = strconv.Itoa(date.Year())
	dateEntries[1] = strconv.Itoa(int(date.Month())); if len(dateEntries[1]) == 1 { dateEntries[1] = "0" + dateEntries[1] }
	dateEntries[2] = strconv.Itoa(date.Day()); if len(dateEntries[2]) == 1 { dateEntries[2] = "0" + dateEntries[2] }
	return strings.Join(dateEntries, "-")
}
