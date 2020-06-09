// Tool to update test data files.  Real URLs are derived from test data file names and urlmap.yaml

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
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

func getDataFromExt(f *TestDataEntryInternal) {
	if len(f.ExtURL) == 0 {
		return
	}
	var resp *http.Response
	var err error
	if strings.ToUpper(f.Method) == "GET" {
		resp, err = http.Get(f.ExtURL)
	} else if strings.ToUpper(f.Method) == "POST" {
		if len(f.ReqFile) == 0 {
			log.Printf("Error ReqFile is missing, url %v", f.ExtURL)
			return
		}
		var b []byte
		b, err = ioutil.ReadFile(f.ReqFile)
		if err != nil {
			log.Printf("Error ReqFile could not be read, error %v, url %v", err.Error(), f.ReqFile)
			return
		}
		resp, err = http.Post(f.ExtURL, "application/json", bytes.NewBuffer(b))
	}
	if err != nil {
		log.Printf("Error reading from external service, err %v, url %v", err.Error(), f.ExtURL)
		return
	}
	if resp.StatusCode != 200 {
		log.Printf("Non-OK status from external service, status %v, url %v", resp.StatusCode, f.ExtURL)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading from external service, err %v, url %v", err.Error(), f.ExtURL)
		return
	}
	if len(body) == 0 {
		log.Printf("Empty response from external service, url %v", f.ExtURL)
		return
	}
	// write to file
	outFile := f.Filename
	err = ioutil.WriteFile(outFile, body, 0644)
	if err != nil {
		log.Printf("Could not write response to file, err %v, file %v", err.Error(), outFile)
		return
	}
	fmt.Printf("Response file written, %v bytes, url %v, file %v\n", len(body), f.ExtURL, outFile)
}

func usage() {
	fmt.Println("mockserver           -- start mock server")
	fmt.Println("mockserver -update   -- update data files from external services")
}

func doStartServer(basedir string) {
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

func doUpdate(basedir string) {
	if err := readFileList(basedir + "/mock"); err != nil {
		log.Fatalf("Could not read data file list, err %v", err.Error())
		return
	}

	for _, f := range files {
		if len(f.ExtURL) >= 0 {
			getDataFromExt(&f)
		}
	}
}

func main() {
	basedir := "."

	nArgs := len(os.Args)
	if nArgs >= 2 {
		if os.Args[1] == "-update" {
			doUpdate(basedir)
			return
		}

		usage()
		return
	}

	doStartServer(basedir)
}
