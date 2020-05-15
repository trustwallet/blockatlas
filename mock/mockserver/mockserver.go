// Tool to update test data files.  Real URLs are derived from test data file names and urlmap.yaml

package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type TestDataEntry struct {
	Filename   string `yaml:"file"`
	MockURL    string `yaml:"mockURL"`
	ExtURL     string `yaml:"extURL,omitempty"`
	ReqFile    string `yaml:"reqFile,omitempty"`
}

type TestDataEntryInternal struct {
	Filename   string   `yaml:"file"`
	MockURL    string   `yaml:"mockURL"`
	ExtURL     string   `yaml:"extURL,omitempty"`
	ParsedURL  *url.URL `yaml:"-"`
	Method     string   `yaml:"-"`
	ReqFile    string   `yaml:"reqFile,omitempty"`
}

var files []TestDataEntryInternal

const extension string = ".json"

// matchQueryParams compares HTTP GET params, checks that all provided params contain all parameters from the expected params, with the same values
func matchQueryParams(expected, actual string) bool {
	if len(expected) == 0 {
		return true
	}
	valuesExp, err := url.ParseQuery(expected)
	if err != nil {
		return false
	}
	valuesAct, err := url.ParseQuery(actual)
	if err != nil {
		return false
	}
	for vv := range valuesExp {
		if _, ok := valuesAct[vv]; !ok {
			// param vv from valuesExp is not present in valuesAct
			return false
		}
		if valuesAct[vv][0] != valuesExp[vv][0] {
			// param is present in both, but value is different
			return false
		}
	}
	// all present with same values
	return true
}

func readFileList(directory string) error {
	filename := directory + "/datafiles.yaml"
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not read index err %v file %v", err.Error(), filename)
		return errors.New("Could not read index file, err " + err.Error() + " file " + filename)
	}
	files1 := []TestDataEntry{}
	err = yaml.Unmarshal(yamlFile, &files1)
	if err != nil {
		log.Fatalf("Could not read index err %v file %v", err.Error(), filename)
		return errors.New("Could not read index file, err " + err.Error() + " file " + filename)
	}

	// some preprocessing
	files = []TestDataEntryInternal{}
	for _, e := range files1 {
		parsedURL, err := url.Parse(e.MockURL)
		if err == nil {
			files = append(files, TestDataEntryInternal{e.Filename, e.MockURL, e.ExtURL, parsedURL, "?", e.ReqFile})
		}
	}
	fmt.Printf("Info about %v data files read\n", len(files))
	return nil
}

func findFileForMockURL(mockURL, queryParams string) (TestDataEntryInternal, error) {
	lasterr := ""
	for _, ff := range files {
		//if mockURL[:7] == ff.MockURL[:7] {
		//	fmt.Println("  ", ff.MockURL, mockURL)
		//}
		// simple check
		if mockURL == ff.MockURL {
			return ff, nil
		}
		// check with query params
		if mockURL == ff.ParsedURL.Path {
			if !matchQueryParams(ff.ParsedURL.RawQuery, queryParams) {
				// remember error, but continue trying
				lasterr = "Mismatch in query params, expected " + ff.ParsedURL.RawQuery + ", actual " + queryParams
			} else {
				return ff, nil
			}
		}

	}
	return TestDataEntryInternal{}, errors.New("Could not find matching entry for URL, " + lasterr)
}

func requestHandlerIntern(w http.ResponseWriter, r *http.Request, basedir string) error {
	if r.URL.Path == "/mock/mock-healtcheck" {
		fmt.Fprintf(w, "{\"status\": true, \"msg\": \"Mockserver is alive\"}")
		return nil
	}

	mockURL := r.URL.Path
	if len(mockURL) >= 1 && mockURL[0] == '/' {
		mockURL = mockURL[1:]
	}

	entry, err := findFileForMockURL(mockURL, r.URL.RawQuery)
	if err != nil {
		return err
	}
	// read and return response
	b, err := ioutil.ReadFile(basedir + "/" + entry.Filename)
	if err != nil {
		return errors.New("Could not read data file for request")
	}
	fmt.Fprintf(w, string(b))
	return nil
}

func requestHandler(w http.ResponseWriter, r *http.Request, basedir string) {
	err := requestHandlerIntern(w, r, basedir)
	if err == nil {
		log.Println("Request ok", r.Method, r.URL.Path)
		return
	}
	// error
	errorMsg := err.Error()
	log.Println("ERROR for request:", errorMsg, r.Method, r.URL.Path)
	fmt.Fprintf(w, "{\"error\": \"" + errorMsg + "\", \"url\": \"" + r.URL.Path + "\"")
}

func main() {
	basedir := "."
	if err := readFileList(basedir + "/mock"); err != nil {
		log.Fatalf("Could not read data file list, err %v", err.Error())
		return
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, basedir)
	})

	port := 3347
	log.Printf("About to listening on port %v", port)
	err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
	if err != nil {
		log.Fatalf("Could not listen on port %v", port)
	}
}
