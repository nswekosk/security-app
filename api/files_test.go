package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDirs = []string{
	"testDir1",
	"testDir2",
}

var testFiles = []struct {
	path     string
	contents string
}{
	{"testDir1/testFile1.txt", "This is test file 1."},
	{"testDir2/testFile2.txt", "This is test file 2."},
}

func TestFileHandlerSuccess(t *testing.T) {
	setUp()
	// Create a request with no path parameter
	req, err := http.NewRequest("GET", "/files", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fileHandler)
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusOK, rr.Code, "invalid status code: (%v)", rr.Code)

	// Create a request with a valid path parameter with test directory created in TestMain
	req, err = http.NewRequest("GET", "/files?path=testDir1", nil)
	assert.Nil(t, err)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	// Assert HTTP Status code 200
	assert.Equalf(t, http.StatusOK, rr.Code, "invalid status code: (%v)", rr.Code)
	assert.Equal(t, `{"name":"testDir1","type":"dir","size":0,"contents":[{"name":"testFile1.txt","type":"file","size":20}]}`, rr.Body.String(), "invalid contents in response")

	req, _ = http.NewRequest("GET", "/files?path=./testDir2", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusOK, rr.Code, "invalid status code: (%v)", rr.Code)
	assert.Equal(t, `{"name":"testDir2","type":"dir","size":0,"contents":[{"name":"testFile2.txt","type":"file","size":20}]}`, rr.Body.String(), "invalid contents in response")
	tearDown()

	// confirm directories are torn down
	req, _ = http.NewRequest("GET", "/files?path=testDir1", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusBadRequest, rr.Code, "invalid status code: (%v)", rr.Code)

	req, _ = http.NewRequest("GET", "/files?path=testDir2", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusBadRequest, rr.Code, "invalid status code: (%v)", rr.Code)
}

func TestFileHandlerNonExistingPath(t *testing.T) {
	setUp()
	req, _ := http.NewRequest("GET", "/files?path=./testDir3", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fileHandler)
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusBadRequest, rr.Code, "invalid status code: (%v)", rr.Code)
	assert.Containsf(t, rr.Body.String(), "You must provide a valid directory path. Path:", "Should contain error message that user must provide a valid directory path")
	tearDown()
}

func TestFileHandlerFilePath(t *testing.T) {
	setUp()
	req, _ := http.NewRequest("GET", "/files?path=./testDir1/testFile1.txt", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fileHandler)
	handler.ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusBadRequest, rr.Code, "invalid status code: (%v)", rr.Code)
	assert.Containsf(t, rr.Body.String(), "You must provide a valid directory path. Path:", "Should contain error message that user must provide a valid directory path")
	tearDown()
}

// Create test files and directories
func setUp() {
	// Create test directories
	for _, dir := range testDirs {
		os.Mkdir(dir, 0755)
	}

	// Create test files
	for _, file := range testFiles {
		err := os.WriteFile(file.path, []byte(file.contents), 0666)
		if err != nil {
			log.Printf("[WARNING] Error writing test file at path (%v): (%v)", file.path, err)
		}
	}
}

// Tear down test files and directories
func tearDown() {
	for _, file := range testFiles {
		err := os.Remove(file.path)
		if err != nil {
			log.Printf("[WARNING] Error tearing down files at test file at path (%v): (%v)", file.path, err)
		}
	}

	for _, dir := range testDirs {
		err := os.Remove(dir)
		if err != nil {
			log.Printf("[WARNING] Error tearing down directories at test file at path (%v): (%v)", dir, err)
		}
	}
}
