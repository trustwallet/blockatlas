// Tool to update test data files.  Real URLs are derived from test data file names and index.yaml

package main

import (
	"errors"
//	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type UrlEntry struct {
	Url string
}

var urlList map[string]UrlEntry

func enumerateDataFiles(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	const extension string = ".json"
	filenames := make([]string, 0)
	for _, f := range files {
		if len(f.Name()) > 0 {
			filename := f.Name()
			if strings.HasSuffix(filename, extension) {
				filename = filename[:len(filename)-len(extension)]
				filenames = append(filenames, filename)
			}
		}
	}
	return filenames
}

func readUrlList() (map[string]UrlEntry, error) {
	filename := "index.yaml"
	list := map[string]UrlEntry{}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return list, errors.New("Could not read index file")
	}
	err = yaml.Unmarshal(yamlFile, list)
	if err != nil {
		return list, errors.New("Could not read index file")
	}
	return list, nil
}

func getRealUrl(mockUrl string) string {
	for mockUrlRoot := range urlList {
		if strings.HasPrefix(mockUrl, mockUrlRoot) {
			realUrlRoot := urlList[mockUrlRoot].Url
			realUrl := strings.Replace(mockUrl, mockUrlRoot, realUrlRoot, -1)
			return realUrl
		}
	}
	return ""
}

func processFile(escapedMockUrl, httpMethod, folder string) {
	log.Printf("Processing %v", escapedMockUrl)
	mockUrl, err := url.QueryUnescape(escapedMockUrl)
	if err == nil {
		realUrl := getRealUrl(mockUrl)
		if len(realUrl) == 0 {
			log.Printf("Could not obtain real URL")
			return
		}
		if httpMethod == "GET" {
			resp, err := http.Get(realUrl)
			if err != nil {
				log.Printf("Error reading from external service, err %v, url %v", err.Error(), realUrl)
				return
			}
			if (resp.StatusCode != 200) {
				log.Printf("Non-OK status from external service, status %v, url %v", resp.StatusCode, realUrl)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading from external service, err %v, url %v", err.Error(), realUrl)
				return
			}
			if len(body) == 0 {
				log.Printf("Empty response from external service, url %v", realUrl)
				return
			}
			// write to file
			outFile := folder + "/" + escapedMockUrl + ".json"
			os.Rename(outFile, outFile + ".bak")
			err = ioutil.WriteFile(outFile, body, 0644)
			if err != nil {
				log.Printf("Could not write response to file, err %v, file %v", err.Error(), outFile)
				return
			}
			log.Printf("Response file written, %v bytes, url %v, file %v", len(body), realUrl, outFile)
			return
		}
	}
}

func main() {
	urlList1, err := readUrlList()
	if err != nil {
		log.Fatal(err)
		return
	}
	urlList = urlList1

	log.Printf("Enumerating test data files ... ")
	files := enumerateDataFiles("./get")
	log.Printf("%v data files found in get", len(files))
	for _, f := range files {
		processFile(f, "GET", "./get")
	}
	files = enumerateDataFiles("./post")
	log.Printf("%v data files found in post", len(files))
	for _, f := range files {
		processFile(f, "POST", "./post")
	}
}
