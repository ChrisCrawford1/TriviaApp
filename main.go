package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/chriscrawford1/trivia_app/helpers"
)

var baseURL string = "https://opentdb.com/api.php?"

type question struct {
	category      string `json:"category"`
	difficulty    string `json:"difficulty"`
	correctAnswer string `json:"correct_answer"`
}

func main() {
	questionAmount := flag.Int("questions", 10, "The number of questions you wish to get")
	difficulty := flag.String("difficulty", "easy", "The difficulty of the questions returned")
	flag.Parse()

	params := make(map[string]string)
	params["amount"] = strconv.Itoa(*questionAmount)
	params["difficulty"] = *difficulty

	query := helpers.BuildQuery(baseURL, params)
	sendQuery(query)
}

func sendQuery(query string) {
	res, _ := http.Get(query)
	data, _ := ioutil.ReadAll(res.Body)
	_ = data
}
