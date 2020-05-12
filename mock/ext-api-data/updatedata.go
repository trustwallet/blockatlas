// Tool to update test data files.  Real URLs are derived from test data file names and urlmap.yaml

package main

import (
	"bufio"
	"errors"
	//"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

type UrlEntry struct {
	Url string
}

var urlList map[string]UrlEntry

// Enumerate data files (extension .json) in given directoty; return filenames without extension == escaped relative URL paths
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
	filename := "urlmap.yaml"
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

// Normalize a relative path URL.  If there are query parameters, they are sorted.
func normalizeUrl(inurl string) string {
	parsedUrl, err := url.Parse(inurl)
	if err != nil {
		return inurl
	}
	if len(parsedUrl.RawQuery) == 0 {
		// no query, nothing to sort
		return inurl
	}
	values, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return inurl
	}
	if len(values) <= 1 {
		// 1 or less, nothing to sort
		return inurl
	}
	// sort values by key
    var keys []string
    for k := range values {
        keys = append(keys, k)
    }
    sort.Strings(keys)

	var querySorted string = ""
	idx := 0
    for _, k := range keys {
		if idx > 0 {
			querySorted += "&"
		}
		querySorted += k + "=" + values[k][0]
		idx++
	}
	
	// build URL
	parsedUrl.RawQuery = querySorted
	normalized := parsedUrl.String()
	return normalized
}

// Given a mock URL relative path, return corresponding full path of the external service.  Mapping from urlList is used.
func getRealUrl(mockUrl string) string {
	for mockUrlRoot := range urlList {
		if strings.HasPrefix(mockUrl, mockUrlRoot) {
			realUrlRoot := urlList[mockUrlRoot].Url
			realUrl := strings.Replace(mockUrl, mockUrlRoot, realUrlRoot, -1)
			return realUrl
		}
	}
	// none found
	return ""
}

// Process a test data file: given the esacaped relative path (== filename without extension), get the real URL, and retrieve response from there.
// The old file is renamed to .bak (unless there is error), and the result is written to the old name.
func processFile(escapedMockUrl, httpMethod, folder string) {
	log.Printf("Processing %v", escapedMockUrl)
	mockUrl, err := url.QueryUnescape(escapedMockUrl)
	if err != nil {
		log.Printf("Could not un-escape url, err %v, url %v", err.Error(), escapedMockUrl)
		return
	}
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
		err = os.Rename(outFile, outFile + ".bak")
		if err != nil {
			log.Printf("Rename to Bak failed, err %v, file %v", err.Error(), outFile)
			return
		}
		err = ioutil.WriteFile(outFile, body, 0644)
		if err != nil {
			log.Printf("Could not write response to file, err %v, file %v", err.Error(), outFile)
			return
		}
		log.Printf("Response file written, %v bytes, url %v, file %v", len(body), realUrl, outFile)
		return
	}
}

func processInitial(initialMockUrlListFileName, folder string) {
	file, err := os.Open(initialMockUrlListFileName)
    if err != nil {
		log.Fatal(err)
		return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		url0 := scanner.Text()
		nurl := normalizeUrl(url0)
		fn := folder + "/" + url.QueryEscape(nurl) + ".json"
		log.Println(fn)
		err = ioutil.WriteFile(fn, []byte{}, 0644)
		if err != nil {
			log.Printf("Could not write empty data file, err %v, file %v", err.Error(), fn)
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main() {
	/*
	fmt.Println(normalizeUrl("mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1"))
	fmt.Println(normalizeUrl("mock/kava-api/txs?page=1&message.sender=kava1l8va&limit=25"))
	fmt.Println(normalizeUrl("mock/kava-api/txs?message.sender=kava1l8va&page=1&limit=25"))

	processInitial("out", "./get")
	return
	*/

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
