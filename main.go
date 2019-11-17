package main

import (
	"flag"
	"fmt"
	"strconv"
)

var baseURL string = "https://opentdb.com/api.php?"

//	"github.com/chriscrawford1/trivia_app/helpers"
// // https://opentdb.com/api.php?amount=10&difficulty=easy

type question struct {
	category      string
	difficulty    string
	correctAnswer string
}

func main() {
	questionAmount := flag.Int("questions", 10, "The number of questions you wish to get")
	difficulty := flag.String("difficulty", "easy", "The difficulty of the questions returned")
	flag.Parse()

	params := make(map[string]string)
	params["questionAmount"] = strconv.Itoa(*questionAmount)
	params["difficulty"] = *difficulty

	buildQuery(params)
}

func buildQuery(params map[string]string) {
	var count int
	for i, line := range params {
		count++

		fmt.Printf("Index [%s] contains [%s]\n", i, line)
	}
}
