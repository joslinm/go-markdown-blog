// +build main

package main

import (
	"github.com/madari/goskirt"
	"net/http"
	"io/ioutil"
	"fmt"
)
func rootHandler(w http.ResponseWriter, request *http.Request) {
	files, _ := ioutil.ReadDir("./posts")
	data := ""
	for _,value := range files {
		// Ignore hidden files
		if value.Name()[0] == '.' {
			continue
		}
		data += value.Name() + "\n============="
		file_data, err := ioutil.ReadFile("./posts/" + value.Name())
		if (err != nil) {
			fmt.Printf("\nGot error!\n%s\n", err)
		}
		fmt.Printf("\nGot data: %s\n", string(file_data[:]))
		data += "\n" + string(file_data[:])
		fmt.Printf("\nOverall data: %s\n", data)
	}
	skirt := goskirt.Goskirt {
		goskirt.EXT_AUTOLINK | goskirt.EXT_STRIKETHROUGH,
		goskirt.HTML_SMARTYPANTS | goskirt.HTML_USE_XHTML,
	}

	skirt.WriteHTML(w, []byte(data))
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

