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
