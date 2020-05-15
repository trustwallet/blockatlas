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
)

type TestDataEntry struct {
	Filename   string `yaml:"file"`
	MockURL    string `yaml:"mockURL"`
}

type TestDataEntryInternal struct {
	Filename   string
	MockURL    string
	ParsedURL  *url.URL
	Method     string
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
	filename := directory + "/files.yaml"
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
			files = append(files, TestDataEntryInternal{e.Filename, e.MockURL, parsedURL, "?"})
		}
	}
	fmt.Printf("Info about %v data files read\n", len(files))
	return nil
}

func findFileForMockURL(mockURL, queryParams string) (TestDataEntryInternal, bool) {
	for _, ff := range files {
		//if mockURL[:7] == ff.MockURL[:7] {
		//	fmt.Println("  ", ff.MockURL, mockURL)
		//}
		// simple check
		if mockURL == ff.MockURL {
			return ff, true
		}
		// check with query params
		if mockURL == ff.ParsedURL.Path {
			if !matchQueryParams(ff.ParsedURL.RawQuery, queryParams) {
				log.Printf("Mismatch in query params, expected %v, actual %v", ff.ParsedURL.RawQuery, queryParams)
			} else {
				return ff, true
			}
		}

	}
	return TestDataEntryInternal{}, false
}

func requestHandler(w http.ResponseWriter, r *http.Request, basedir string) {
	log.Println(r.Method, r.URL.Path, r.URL.RawPath, r.URL.RawQuery)
	if r.URL.Path == "/mock/mock-healtcheck" {
		fmt.Fprintf(w, "{\"status\": true, \"msg\": \"Mockserver is alive\"}")
		return
	}

	mockURL := r.URL.Path
	if len(mockURL) >= 1 && mockURL[0] == '/' {
		mockURL = mockURL[1:]
	}

	entry, ok := findFileForMockURL(mockURL, r.URL.RawQuery)
	if !ok {
		log.Printf("Can't handle mock URL %v!", mockURL)
		fmt.Fprintf(w, "{\"error\": \"Can't handle mock URL\"}")
		return
	}
	// read and return response
	b, err := ioutil.ReadFile(basedir + "/" + entry.Filename)
	if err != nil {
		log.Fatalf("Can't read data file file %v url %v!", entry.Filename, mockURL)
		fmt.Fprintf(w, "{\"error\": \"Can't read file\"}")
		return
	}
	fmt.Fprintf(w, string(b))
}

func main() {
	basedir := "."
	//loadFiles(basedir)
	if err := readFileList(basedir); err != nil {
		log.Fatalf("Could not read data file list, err %v", err.Error())
		return
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, basedir)
	})

    http.ListenAndServe(":3347", nil)
}
