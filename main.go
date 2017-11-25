package main

import "flag"
import "fmt"
import "os"
import "strings"
import "io/ioutil"
import "net/http"
import "encoding/json"

var movieName = flag.String("movie", "", "movie name")
var keyPath   = flag.String("key", "key.txt", "API key file path")
var BASEURL   = "http://www.omdbapi.com/?t=%s&apikey=%s"

type omdbMovie struct {
	Title    string `json:"Title"`
	Director string `json:"Director"`
	Poster   string `json:"Poster"`
}

func parseFlags() bool {
	flag.Parse()
	if len(*movieName) == 0 {
		return false
	}
	if len(*keyPath) == 0 {
		return false
	}
	return true
}

func readApiKey(filepath string) (string, error) {
	apiKeyBytes, err := ioutil.ReadFile(filepath)
	apiKey := string(apiKeyBytes[:])
	apiKey = strings.TrimSpace(apiKey)
	return apiKey, err
}

func getMovieData(movieName, apiKey string) (*omdbMovie, error) {
	movieName = strings.Replace(movieName, " ", "+", -1)
	url := fmt.Sprintf(BASEURL, movieName, apiKey)

	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	jb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(omdbMovie)
	if err = json.Unmarshal(jb, data); err != nil {
		return nil, err
	}
	return data, nil
}

func getPoster(movie *omdbMovie) {
	posterURL := movie.Poster
	fmt.Println(posterURL)
}

func main() {
	if !parseFlags() {
		flag.PrintDefaults()
		os.Exit(1)
	}
	apiKey, err := readApiKey(*keyPath)
	if err != nil {
		panic(err)
	}
	movie, err := getMovieData(*movieName, apiKey)
	if err != nil {
		panic(err)
	}
	getPoster(movie)
}
