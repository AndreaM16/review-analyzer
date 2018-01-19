package main

import (
	"github.com/andream16/review-analyzer/algorithm"
	"fmt"
	"os"
)

func main() {
	err := algorithm.StartAlgorithm(); if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}