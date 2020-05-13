// Tool to update test data files.  Real URLs are derived from test data file names and urlmap.yaml

package main

import (
	//"bufio"
	//"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type URLEntry struct {
	URL string
}

var urlList map[string]URLEntry

type TestDataEntry struct {
	Filename   string
	Subdir     string
	HTTPMethod string
}

const extension string = ".json"

// Enumerate data files (extension .json) in given directoty; return filenames without extension == escaped relative URL paths
func enumerateDataFiles(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	filenames := make([]string, 0)
	for _, f := range files {
		if len(f.Name()) > 0 {
			if strings.HasSuffix(f.Name(), extension) {
				filenames = append(filenames, f.Name())
			}
		}
	}
	return filenames
}

func enumerateAllDataFiles(directory string) []TestDataEntry {
	fmt.Printf("Enumerating test data files ... ")
	subdir := "get"
	filesGet := enumerateDataFiles(directory + "/" + subdir)
	var dataFiles []TestDataEntry
	for _, f := range filesGet {
		dataFiles = append(dataFiles, TestDataEntry{f, subdir, "GET"})
	}
	subdir = "./post"
	filesPost := enumerateDataFiles(directory + "/" + subdir)
	for _, f := range filesPost {
		dataFiles = append(dataFiles, TestDataEntry{f, subdir, "POST"})
	}
	fmt.Printf("  %v data files found\n", len(dataFiles))
	return dataFiles
}

func readURLList(directory string) bool {
	filename := directory + "/urlmap.yaml"
	list := map[string]URLEntry{}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not read index file")
		return false
	}
	err = yaml.Unmarshal(yamlFile, list)
	if err != nil {
		log.Fatal("Could not read index file")
		return false
	}
	urlList = list
	return true
}

// Normalize a relative path URL.  If there are query parameters, they are sorted.
func normalizeURL(inurl string) string {
	parsedURL, err := url.Parse(inurl)
	if err != nil {
		return inurl
	}
	if len(parsedURL.RawQuery) == 0 {
		// no query, nothing to sort
		return inurl
	}
	values, err := url.ParseQuery(parsedURL.RawQuery)
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
	parsedURL.RawQuery = querySorted
	normalized := parsedURL.String()
	return normalized
}

// Given a mock URL relative path, return corresponding full path of the external service.  Mapping from urlList is used.
func getRealURL(mockURL string) string {
	for mockURLRoot := range urlList {
		if strings.HasPrefix(mockURL, mockURLRoot) {
			realURLRoot := urlList[mockURLRoot].URL
			realURL := strings.Replace(mockURL, mockURLRoot, realURLRoot, -1)
			return realURL
		}
	}
	// none found
	return ""
}

// Process a test data file: given the esacaped relative path (== filename without extension), get the real URL, and retrieve response from there.
// The old file is renamed to .bak (unless there is error), and the result is written to the old name.
func processFile(file TestDataEntry, directory string, listOnly bool) {
	fmt.Printf("Filename:   %v\n", file.Filename)
	if !strings.HasSuffix(file.Filename, extension) {
		return
	}
	escapedMockURL := file.Filename[:len(file.Filename)-len(extension)]
	mockURL, err := url.QueryUnescape(escapedMockURL)
	fmt.Printf("Mock URL:   %v\n", mockURL)
	if err != nil {
		log.Printf("Could not un-escape url, err %v, url %v", err.Error(), escapedMockURL)
		return
	}
	realURL := getRealURL(mockURL)
	if len(realURL) == 0 {
		log.Printf("Could not obtain real URL")
		return
	}
	fmt.Printf("Real URL:   %v\n", realURL)
	if listOnly {
		return
	}
	// continue with processing
	if file.HTTPMethod == "GET" {
		resp, err := http.Get(realURL)
		if err != nil {
			log.Printf("Error reading from external service, err %v, url %v", err.Error(), realURL)
			return
		}
		if resp.StatusCode != 200 {
			log.Printf("Non-OK status from external service, status %v, url %v", resp.StatusCode, realURL)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading from external service, err %v, url %v", err.Error(), realURL)
			return
		}
		if len(body) == 0 {
			log.Printf("Empty response from external service, url %v", realURL)
			return
		}
		// write to file
		outFile := directory + "/" + escapedMockURL + ".json"
		err = os.Rename(outFile, outFile+".bak")
		if err != nil {
			log.Printf("Rename to Bak failed, err %v, file %v", err.Error(), outFile)
			return
		}
		err = ioutil.WriteFile(outFile, body, 0644)
		if err != nil {
			log.Printf("Could not write response to file, err %v, file %v", err.Error(), outFile)
			return
		}
		fmt.Printf("Response file written, %v bytes, url %v, file %v", len(body), realURL, outFile)
		return
	}
}

/*
func processInitial(initialMockURLListFileName, folder string) {
	file, err := os.Open(initialMockURLListFileName)
    if err != nil {
		log.Fatal(err)
		return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		url0 := scanner.Text()
		nurl := normalizeURL(url0)
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
*/

func PrintUsage() {
	fmt.Println("Usage:")
	prog := "testdatatool"
	fmt.Printf("  %v list <folder>        List all test data files.                        Ex.: %v list .\n", prog, prog)
	fmt.Printf("  %v update <datafile>    Update a test data from external source.         Ex.: %v update get/mock2Fsomthing.json\n", prog, prog)
	fmt.Printf("  %v updateall <folder>   Update all test data files from external source. Ex.: %v updateall .\n", prog, prog)
	fmt.Printf("  %v help                 Print this help\n", prog)
	fmt.Printf("Mapping to real URLs is taken from file urlmap.yaml\n")
}

// ListAll lists all test data files and prints info about them
func ListAll(directory string) {
	dataFiles := enumerateAllDataFiles(directory)
	i := 0
	for _, f := range dataFiles {
		i++
		fmt.Printf("%v:\n", i)
		processFile(f, directory, true)
	}
}

func httpMethodFromFilename(filename string) (method, lastDir string, ok bool) {
	lastDir = filepath.Base(filepath.Dir(filename))
	method = strings.ToUpper(lastDir)
	if method == "GET" || method == "POST" {
		return method, lastDir, true
	}
	log.Printf("Could not determine HTTP method for file %v", filename)
	return "", lastDir, false
}

func UpdateFile(filename string) {
	method, lastDir, ok := httpMethodFromFilename(filename)
	if !ok {
		return
	}
	entry := TestDataEntry{filepath.Base(filename), lastDir, method}
	processFile(entry, filepath.Dir(filename), false)
}

func UpdateAllFiles(directory string) {
	dataFiles := enumerateAllDataFiles(directory)
	i := 0
	for _, f := range dataFiles {
		i++
		fmt.Printf("%v:\n", i)
		processFile(f, directory, false)
	}
}

func main() {
	nArgs := len(os.Args)
	if nArgs <= 1 {
		// no args, list by default
		dir := "."
		if !readURLList(dir) {
			return
		}
		ListAll(dir)
		return
	}

	switch os.Args[1] {
	case "list":
		if nArgs < 3 {
			PrintUsage()
			return
		}
		dir := os.Args[2]
		if !readURLList(dir) {
			return
		}
		ListAll(dir)
		return

	case "update":
		if nArgs < 3 {
			PrintUsage()
			return
		}
		filename := os.Args[2]
		if !readURLList(".") {
			return
		}
		UpdateFile(filename)
		return

	case "updateall":
		if nArgs < 3 {
			PrintUsage()
			return
		}
		dir := os.Args[2]
		if !readURLList(".") {
			return
		}
		UpdateAllFiles(dir)
		return

	case "help":
		PrintUsage()
		return

	default:
		PrintUsage()
		return
	}
}
