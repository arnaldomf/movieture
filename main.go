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

func requestMoviePicture(movieName, apiKey string) (string, error) {
	movieName = strings.Replace(movieName, " ", "+", -1)
	url := fmt.Sprintf(BASEURL, movieName, apiKey)

	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if resp.Body == nil {
		return "", nil
	}
	defer resp.Body.Close()
	jb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := make(map[string]interface{})
	if err = json.Unmarshal(jb, &data) {
		return "", err
	}
	fmt.Println(data)
	return , nil
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
	requestMoviePicture(*movieName, apiKey)
}
