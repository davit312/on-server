package fs

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Handler simple struct to handle files list from main function
type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Replace(r.RequestURI, URLRoot, "", 1)
	filePath := filesRoot + "/" + fileName

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`"` + fileName + `"` + ` not found`))
		return
	}

	if fileInfo.IsDir() {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")

		files, err := ioutil.ReadDir(filePath)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`Server error`))
			return
		}

		var page string

		page = top +
			makeList(URLRoot+fileName, files) +
			bottom

		w.Write([]byte(page))
	} else {
		http.ServeFile(w, r, filePath)
	}

}