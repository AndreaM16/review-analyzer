package exec

import (
	"os/exec"
	"strings"
	"strconv"
)

var argv = [2]string{"python", "sentiment/sentiment.py"}

// Executes the python script.
func GetSentimentAnalysisFromSentence(sentence string) (float64, error) {
	out, err := func() ([]byte, error) {
		return func() *exec.Cmd {
			return exec.Command(argv[0], argv[1], sentence)
		}().Output()
	}(); if err != nil {
		return 0.0, err
	}
	parsedResult, parsingError := getFloat64FromString(strings.TrimSuffix(string(out), "\n"))
	if parsingError != nil {
		return 0.0, parsingError
	}
	return parsedResult, nil
}

func getFloat64FromString(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}
