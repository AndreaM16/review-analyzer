package exec

import "os/exec"

var argv = [2]string{"python", "/sentiment/sentiment.py"}

// Executes the python script.
func GetSentimentAnalysisFromSentence(sentence string) (string, error) {
	out, err := func() ([]byte, error) {
		return func() *exec.Cmd {
			return exec.Command(argv[0], argv[1], sentence)
		}().Output()
	}(); if err != nil {
		return "", err
	}
	return	string(out), nil
}
