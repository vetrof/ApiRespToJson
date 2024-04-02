package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	var url string
	fmt.Println("api url: ")
	fmt.Scanf("%s\n", &url)

	// URL API
	// url := "https://api.sampleapis.com/beers/ale"

	// get
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// check ctatus
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "error: %s\n", resp.Status)
		os.Exit(1)
	}

	// decode JSON
	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Fprintf(os.Stderr, "error decode JSON: %v\n", err)
		os.Exit(1)
	}

	// format JSON
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error format json JSON: %v\n", err)
		os.Exit(1)
	}

	// json to terminal
	fmt.Println(string(prettyJSON))

	// current date
	currentTime := time.Now()
	dateTimeStr := currentTime.Format("2006-01-02_15-04-05")

	// filename
	fileName := fmt.Sprintf("json_out_%s.json", dateTimeStr)

	// save to file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error crete file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(prettyJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error record to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("done output.json")
}
