// Tool to update test data files.  Real URLs are derived from test data file names and urlmap.yaml

package main

import (
	//"bufio"
	"bytes"
	"encoding/json"
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
	"strconv"
	"strings"
)

type URLEntry struct {
	URL string
}

type URLMap map[string]URLEntry

var urlList URLMap

type TestDataEntry struct {
	Filename   string
	Basedir    string
	HTTPMethod string
}

const extension string = ".json"
const extensionPost string = ".request_json"

// Enumerate data files (with extension .json) in given directoty; return filenames
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
		dataFiles = append(dataFiles, TestDataEntry{f, directory + "/" + subdir, "GET"})
	}
	subdir = "post"
	filesPost := enumerateDataFiles(directory + "/" + subdir)
	for _, f := range filesPost {
		dataFiles = append(dataFiles, TestDataEntry{f, directory + "/" + subdir, "POST"})
	}
	fmt.Printf("  %v data files found\n", len(dataFiles))
	return dataFiles
}

func readURLList(directory string) bool {
	filename := directory + "/urlmap.yaml"
	list := map[string]URLEntry{}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not read index file %v", filename)
		return false
	}
	err = yaml.Unmarshal(yamlFile, list)
	if err != nil {
		log.Fatalf("Could not read index file %v", filename)
		return false
	}
	urlList = list
	return true
}

func mockURLFromFilename(filename string) (mockURL, counter string, err error) {
	// strip extension
	escapedMockURL := filename
	if (strings.HasSuffix(filename, extension)) {
		escapedMockURL = escapedMockURL[:len(escapedMockURL)-len(extension)]
	}
	// strip counter if present
	counter = ""
	counterExt := filepath.Ext(escapedMockURL)
	if len(counterExt) >= 2 && len(counterExt) <= 5 {
		if _, err := strconv.Atoi(counterExt[1:]); err == nil {
			// there is a counter extension
			counter = counterExt[1:]
			escapedMockURL = escapedMockURL[:len(escapedMockURL)-len("." + counter)]
		}
	}
	mockURL, err = url.QueryUnescape(escapedMockURL)
	if err != nil {
		log.Printf("Could not un-escape url, err %v, url %v", err.Error(), escapedMockURL)
		return mockURL, counter, err
	}
	return mockURL, counter, nil
}

func filenameFromMockURL(mockURL, counter string) string {
	filename := url.QueryEscape(mockURL)
	if len(counter) > 0 {
		filename = filename + "." + counter
	}
	filename = filename + extension
	return filename
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
func getRealURL(mockURL string, urlMap URLMap) string {
	for mockURLRoot := range urlMap {
		if strings.HasPrefix(mockURL, mockURLRoot) {
			realURLRoot := urlMap[mockURLRoot].URL
			realURL := strings.Replace(mockURL, mockURLRoot, realURLRoot, -1)
			return realURL
		}
	}
	// none found
	return ""
}

// Given a real external API URL, return the corressponding mock path.  Mapping must exists.
func getMockURL(realURL string, urlMap URLMap) string {
	for mockURLRoot := range urlMap {
		realURLRoot := urlMap[mockURLRoot].URL
		if strings.HasPrefix(realURL, realURLRoot) {
			mockURL := strings.Replace(realURL, realURLRoot, mockURLRoot, -1)
			return mockURL
		}
	}
	// none found
	return ""

}

func readOrWritePostRequestData(entry TestDataEntry, postRequestData string) string {
	// check if request data file exists
	postFile := strings.Replace(entry.Basedir + "/" + entry.Filename, extension, extensionPost, -1)
	b, err := ioutil.ReadFile(postFile)
	if err == nil {
		// file exists, return its content
		fmt.Printf("Post request data read from file  %v\n", postFile)
		return string(b)
	}
	// file could not be read, use parameter, write it to file
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(postRequestData), "", "\t")
	if err != nil {
		log.Printf("JSON parse error, err %v, data %v\n", err.Error(), postRequestData)
		return postRequestData
	}
	err = ioutil.WriteFile(postFile, prettyJSON.Bytes(), 0644)
	if err != nil {
		log.Printf("Could not write post params to new file, err %v, file %v", err.Error(), postFile)
		return postRequestData
	}
	fmt.Printf("Post request data written to file  %v\n", postFile)
	return postRequestData
}

// Process a test data file: given the esacaped relative path (== filename without extension), get the real URL, and retrieve response from there.
// The old file is renamed to .bak (unless there is error), and the result is written to the old name.
func processFile(file TestDataEntry, listOnly bool, postRequestData string) {
	fmt.Printf("Filename:   %v\n", file.Filename)
	if !strings.HasSuffix(file.Filename, extension) {
		return
	}
	mockURL, _, err := mockURLFromFilename(file.Filename)
	if err != nil {
		log.Printf("Could not obtain URL from filename, err %v, filename %v", err.Error(), file.Filename)
		return
	}
	fmt.Printf("Mock URL:   %v\n", mockURL)
	realURL := getRealURL(mockURL, urlList)
	if len(realURL) == 0 {
		log.Printf("Could not obtain real URL")
		return
	}
	fmt.Printf("Real URL:   %v\n", realURL)

	if listOnly {
		return
	}
	// continue with processing

	var resp *http.Response
	switch (file.HTTPMethod) {
		case "GET":
			resp, err = http.Get(realURL)
		case "POST":
			postData := readOrWritePostRequestData(file, postRequestData)
			resp, err = http.Post(realURL, "application/json", bytes.NewBuffer([]byte(postData)))
		default:
			log.Printf("Invalid method %v, url %v", file.HTTPMethod, realURL)
			return
	}

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
	outFile := file.Basedir + "/" + file.Filename
	if _, err = os.Stat(outFile); err == nil {
		// file exists, rename
		bakFile := outFile+".bak"
		err = os.Rename(outFile, bakFile)
		if err != nil {
			log.Printf("Rename to Bak failed, err %v, file %v", err.Error(), outFile)
			return
		}
	}
	// pretty print
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		log.Printf("JSON parse error, err %v, file %v\n", err.Error(), outFile)
		return
	}
	err = ioutil.WriteFile(outFile, prettyJSON.Bytes(), 0644)
	if err != nil {
		log.Printf("Could not write response to file, err %v, file %v", err.Error(), outFile)
		return
	}
	fmt.Printf("Response file written, %v bytes, url %v, file %v\n", len(body), realURL, outFile)
}

func PrintUsage() {
	fmt.Println("Usage:")
	prog := "testdatatool"
	fmt.Printf("  %v help                   Print this help\n", prog)
	fmt.Printf("  %v list <folder>          List all test data files.                        Ex.: %v list .\n", prog, prog)
	fmt.Printf("  %v update <datafile>      Update a test data from external source.         Ex.: %v update get/mock2Fsomthing.json\n", prog, prog)
	fmt.Printf("  %v updateall <folder>     Update all test data files from external source. Ex.: %v updateall .\n", prog, prog)
	fmt.Printf("  %v add <realURL> <method> [<postReqData>]\n", prog)
	fmt.Printf("                                    Create a new data file from a real full URL.     Ex.: %v add https://api.trongrid.io/v1/assets/1002798 get\n", prog)
	fmt.Printf("                                    Ex.: %v add https://vethor-pubnode.digonchain.com/logs/transfer post '{\"options\":{\"offset\":0,\"limit\":15},...}'\n", prog)
	fmt.Printf("Mapping to real URLs is taken from file urlmap.yaml\n")
}

// ListAll lists all test data files and prints info about them
func ListAll(directory string) {
	dataFiles := enumerateAllDataFiles(directory)
	i := 0
	for _, f := range dataFiles {
		i++
		fmt.Printf("%v:\n", i)
		processFile(f, true, "")
	}
}

func httpMethodFromFilename(filename string) (method string, ok bool) {
	lastDir := filepath.Base(filepath.Dir(filename))
	method = strings.ToUpper(lastDir)
	if method == "GET" || method == "POST" {
		return method, true
	}
	log.Printf("Could not determine HTTP method for file %v", filename)
	return "", false
}

func AddFile(realURL, method, directory, postRequestData string) {
	mockURL := getMockURL(realURL, urlList)
	if len(mockURL) == 0 {
		log.Printf("Could not obtain mock URL for URL %v.  Does mapping exists between real hostname and mock prefix?", realURL)
		return
	}
	subdir := strings.ToLower(method)
	fulldir := directory + "/" + subdir
	counter := ""
	filename := filenameFromMockURL(mockURL, counter)


	if _, err := os.Stat(fulldir + "/" + filename); err == nil {
		// file exists, check different counters
		for i := 1; i < 10000; i++ {
			counter = "0000" + strconv.Itoa(i)
			counter = counter[len(counter)-4:len(counter)]
			fmt.Println(i, counter)
			filename = filenameFromMockURL(mockURL, counter)
			if _, err = os.Stat(fulldir + "/" + filename); err != nil {
				break;
			}
		}
		fmt.Printf("File already exists, chose counter:  %v filename %v\n", counter, filename)
	}

	fmt.Printf("Mock path and filename:  %v  %v\n", mockURL, filename)
	entry := TestDataEntry{filename, fulldir, strings.ToUpper(method)}
	processFile(entry, false, postRequestData)
}

func UpdateFile(filename string) {
	method, ok := httpMethodFromFilename(filename)
	if !ok {
		return
	}
	entry := TestDataEntry{filepath.Base(filename), filepath.Dir(filename), method}
	processFile(entry, false, "")
}

func UpdateAllFiles(directory string) {
	dataFiles := enumerateAllDataFiles(directory)
	i := 0
	for _, f := range dataFiles {
		i++
		fmt.Printf("%v:\n", i)
		processFile(f, false, "")
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
		if nArgs != 3 {
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
		if nArgs != 3 {
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
		if nArgs != 3 {
			PrintUsage()
			return
		}
		dir := os.Args[2]
		if !readURLList(".") {
			return
		}
		UpdateAllFiles(dir)
		return

	case "add":
		if nArgs < 4 || nArgs > 5 {
			PrintUsage()
			return
		}
		dir := "."
		if !readURLList(dir) {
			return
		}
		arg4 := ""
		if nArgs >= 5 {
			arg4 = os.Args[4]
		}
		AddFile(os.Args[2], os.Args[3], dir, arg4)
		return

	case "help":
		PrintUsage()
		return

	default:
		PrintUsage()
		return
	}
}
