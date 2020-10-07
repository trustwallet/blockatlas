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
	"strings"
)

type TestDataEntry struct {
	Filename string `yaml:"file"`
	MockURL  string `yaml:"mockURL"`
	Method   string `yaml:"method"`
	ExtURL   string `yaml:"extURL,omitempty"`
	ReqFile  string `yaml:"reqFile,omitempty"`
	ReqField string `yaml:"reqField,omitempty"`
}

type TestDataEntryInternal struct {
	TestDataEntry
	ParsedURL *url.URL `yaml:"-"`
}

var files []TestDataEntryInternal

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
			files = append(files, TestDataEntryInternal{e, parsedURL})
		}
	}
	fmt.Printf("Info about %v data files read\n", len(files))
	return nil
}

func fieldValueFromJson(json, field string) (string, error) {
	fieldIdx := strings.Index(json, field)
	if fieldIdx < 0 {
		return "", errors.New("Field not found")
	}
	rest := json[fieldIdx+len(field)+1:]
	firstQuote := strings.Index(rest, "\"")
	if fieldIdx < 0 {
		return "", errors.New("1st quote not found")
	}
	rest = rest[firstQuote+1:]
	secondQuote := strings.Index(rest, "\"")
	if fieldIdx < 0 {
		return "", errors.New("1st quote not found")
	}
	val := rest[:secondQuote]
	return val, nil
}

func preprocessRequestJson(j string) string {
	j = strings.Replace(j, "\"", "", -1)
	j = strings.Replace(j, "'", "", -1)
	j = strings.Replace(j, " ", "", -1)
	j = strings.Replace(j, "\n", "", -1)
	j = strings.Replace(j, "\t", "", -1)
	return j
}

func matchRequestDataJson(actualReqData, expReqData, fieldDiscriminator string) error {
	if len(fieldDiscriminator) > 0 {
		valActual, err := fieldValueFromJson(actualReqData, fieldDiscriminator)
		if err != nil {
			return err
		}
		valExp, err := fieldValueFromJson(expReqData, fieldDiscriminator)
		if err != nil {
			return err
		}
		if valExp == valActual {
			// OK, match
			log.Println("Request data match based on discriminator field " + fieldDiscriminator + " val " + valActual)
			return nil
		}
		return errors.New("Request data mismatch, discriminator field " + fieldDiscriminator + " actual " + valActual + " expected " + valExp)
	}
	// no field separator, full  match
	if actualReqData == expReqData {
		return nil
	}
	v1r := preprocessRequestJson(actualReqData)
	v2r := preprocessRequestJson(expReqData)
	if v1r == v2r {
		return nil
	}
	return errors.New("Mismatch in request data, actual '" + actualReqData + "' expected '" + expReqData + "'")
}

func findFileForMockURL(mockURL, method, queryParams, requestBody string) (TestDataEntryInternal, error) {
	lasterr := ""
	for _, ff := range files {
		if method != ff.Method {
			continue
		}
		// simple check
		if mockURL != ff.ParsedURL.Path {
			continue
		}
		// check query params
		if len(queryParams) > 0 {
			if !matchQueryParams(ff.ParsedURL.RawQuery, queryParams) {
				// mismatch in query, remember message, but continue trying
				lasterr = "Mismatch in query params, expected " + ff.ParsedURL.RawQuery + ", actual " + queryParams
				continue
			}
		}
		// check request data
		if len(requestBody) > 0 {
			// read request file
			reqFileB, err := ioutil.ReadFile(ff.ReqFile)
			if err == nil {
				expectedRequestData := string(reqFileB)
				if err = matchRequestDataJson(requestBody, expectedRequestData, ff.ReqField); err != nil {
					// mismatch in request data, remember message, but continue trying
					lasterr = "Mismatch in request data " + err.Error()
					continue
				}
			}
		}
		// all matches
		return ff, nil
	}
	return TestDataEntryInternal{}, errors.New("Could not find matching entry for URL, " + lasterr)
}

func requestHandlerIntern(w http.ResponseWriter, r *http.Request, method, body, basedir string) error {
	mockURL := r.URL.Path
	entry, err := findFileForMockURL(mockURL, method, r.URL.RawQuery, body)
	if err != nil {
		return err
	}
	// read and return response
	b, err := ioutil.ReadFile(basedir + "/" + entry.Filename)
	if err != nil {
		return errors.New("Could not read data file for request")
	}
	fmt.Fprint(w, string(b))
	return nil
}

func requestHandler(w http.ResponseWriter, r *http.Request, basedir string) {
	// read body
	body := ""
	if r.Method == "POST" {
		bodyByte, err := ioutil.ReadAll(r.Body)
		if err == nil {
			body = string(bodyByte)
		}
	}
	err := requestHandlerIntern(w, r, r.Method, body, basedir)
	if err == nil {
		log.Println("Request ok", r.Method, r.URL.Path, r.URL.RawQuery, body)
		return
	}
	// error
	errorMsg := err.Error()
	log.Println("ERROR for request:", errorMsg, r.Method, r.URL.Path, r.URL.RawQuery, body)
	fmt.Fprintf(w, "{\"error\": \""+errorMsg+"\", \"url\": \""+r.URL.Path+"\"")
}

func main() {
	basedir := "."
	if err := readFileList(basedir + "/mock"); err != nil {
		log.Fatalf("Could not read data file list, err %v", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, basedir)
	})

	port := 3347
	log.Printf("About to listen on port %v", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatalf("Could not listen on port %v", port)
	}
}
