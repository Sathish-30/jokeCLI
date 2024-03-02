package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/icelain/jokeapi"
)

type JokesResp struct {
	Error    bool
	Category string
	JokeType string
	Joke     []string
	Flags    map[string]bool
	Id       float64
	Lang     string
}

func main() {
	api := jokeapi.New()
	ctgs := []string{}
	if len(os.Args) >= 1 {
		ctgs = append(ctgs, os.Args[1:]...)
	}
	ctgs = append(ctgs, "Dark")
	api.Set(jokeapi.Params{Categories: ctgs})
	res, err := api.Fetch()
	if err != nil {
		panic("Unable to fetch data")
	}
	var joke strings.Builder
	for _, val := range res.Joke {
		joke.WriteString(fmt.Sprint(val, " "))
	}
	c := color.New(color.FgGreen, color.Bold)
	c.Println(joke.String())
}
