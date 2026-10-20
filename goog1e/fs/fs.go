//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js

package fs

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/rakyll/statik/example/statik"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func FS2main() {
	statikFS, err := fs.New()
	if err != nil {
		logrus.Fatal(err)
	}

	http.Handle("/p/", http.StripPrefix("/p/", http.FileServer(statikFS)))
	var port int = 8994
	http.ListenAndServe(":"+string(port), nil)
}
func FS3main() {
	statikFS, err := fs.New()
	if err != nil {
		logrus.Fatal(err)
	}

	// Access individual files by their paths.
	r, err := statikFS.Open("/c.js")
	if err != nil {
		logrus.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(string(contents))
}
