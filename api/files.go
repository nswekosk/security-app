package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DirInfo is a struct to hold directory info
type DirInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

// fileHandler handles requests for file information
func fileHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	basePath := "./"
	if pathParam, ok := params["path"]; ok && len(pathParam) > 0 {
		basePath = filepath.Join(basePath, filepath.Clean("/"+pathParam[0]))
	}

	dirs, err := os.ReadDir(basePath)
	if err != nil {
		if strings.Contains(err.Error(), "not a directory") || strings.Contains(err.Error(), "no such file or directory") {
			log.Printf("[fileHandler] Error when user provided an invalid directory (%v): (%v)", basePath, err)
			http.Error(w, fmt.Sprintf("You must provide a valid directory path. Path: (%s)", basePath), http.StatusBadRequest)
			return
		}
		log.Printf("[fileHandler] Error reading directory (%v): (%v)", basePath, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dirInfoL := []DirInfo{}
	for _, dir := range dirs {
		dirInfo, err := dir.Info()
		if err != nil {
			log.Printf("[fileHandler] Error getting file info for dir with name (%v): (%v)", dir.Name(), err)
			continue
		}
		dirType := "dir"
		var size int64
		if !dir.IsDir() {
			dirType = "file"
			size = dirInfo.Size()
		}

		dirInfoL = append(dirInfoL, DirInfo{
			Name: dir.Name(),
			Type: dirType,
			Size: size,
		})
	}

	response := FileResponse{
		Name:     basePath,
		Type:     "dir",
		Size:     0,
		Contents: dirInfoL,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("[fileHandler] Error marshaling response: (%v)", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
