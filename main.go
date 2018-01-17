package main

import (
	executor "github.com/andream16/review-analyzer/exec"
	"fmt"
)

func main() {
	sentiment, err := executor.GetSentimentAnalysisFromSentence("Hey there I love you")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("Got %s", sentiment))
	}
}