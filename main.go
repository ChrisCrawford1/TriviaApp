package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/chriscrawford1/trivia_app/helpers"
)

var baseURL string = "https://opentdb.com/api.php?"

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
	params["amount"] = strconv.Itoa(*questionAmount)
	params["difficulty"] = *difficulty

	query := helpers.BuildQuery(baseURL, params)
	fmt.Printf("Query is %s", query)
}
