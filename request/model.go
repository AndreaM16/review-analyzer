package request

type Item struct {
	Item string `json:"item"`
}

type Items struct {
	Items []Item `json:"items"`
}

type Review struct {
	Item string `json:"item,omitempty"`
	Date string `json:"date,omitempty"`
	Content string `json:"content,omitempty"`
	Stars uint `json:"stars,omitempty"`
}

type Reviews struct {
	Item string `json:"item,omitempty"`
	Reviews []Review `json:"reviews"`
}

type AnalyzedReviews struct {
	Reviews []AnalyzedReview `json:"reviews"`
}

type AnalyzedReview struct {
	Content string `json:"content,omitempty"`
	Date string `json:"date"`
	Sentiment float64 `json:"sentiment"`
	Stars float64 `json:"stars"`
}

type Prices struct {
	Item string `json:"item"`
	Prices []Price `json:"prices"`
}

type Price struct {
	Price float64 `json:"Price"`
	Date string `json:"date"`
}
