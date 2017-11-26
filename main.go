package main

import "flag"
import "fmt"
import "os"
import "strings"
import "io/ioutil"
import "net/http"
import "encoding/json"

var movieName   = flag.String("movie", "", "movie name")
var keyPath     = flag.String("key", "key.txt", "API key file path")
var releaseYear = flag.String("year", "", "Year of Release")
var BASEURL     = "http://www.omdbapi.com/?apikey=%s&t=%s"

type omdbMovie struct {
	Title    string `json:"Title"`
	Director string `json:"Director"`
	Poster   string `json:"Poster"`
	Year     string `json:"Year"`
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

func getMovieData(movieName, releaseYear, apiKey string) (*omdbMovie, error) {
	movieName = strings.Replace(movieName, " ", "+", -1)
	url := fmt.Sprintf(BASEURL, apiKey, movieName)
	if releaseYear != "" {
		url = url + fmt.Sprintf("&y=%s", releaseYear)
	}
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode %d", resp.StatusCode)
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

func getPoster(movie *omdbMovie) ([]byte, error) {
	posterURL := movie.Poster
	resp, err := http.Get(posterURL)
	if err != nil {
		return nil,err
	}
	if resp.StatusCode != 200 {
		return nil,fmt.Errorf("StatusCode: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body, nil
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
	movie, err := getMovieData(*movieName, *releaseYear, apiKey)
	if err != nil {
		panic(err)
	}
	_, err = getPoster(movie)
	if err != nil {
		panic(err)
	}
}
