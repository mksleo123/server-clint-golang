package solution

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"../types1"
)

var fileinfos1 []types1.FileInfo

func returnAllFileinfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllFileinfo")
	json.NewEncoder(w).Encode(fileinfos)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/fileinfo", returnAllFileinfo)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func filetype12(c bool) string {
	if c == true {
		return "Dir"
	}
	return "file"
}

func PartD() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		fileStat, err := os.Stat(f.Name())
		if err != nil {
			log.Fatal(err)
		}

		fileinffo := &types1.FileInfo{
			Filepath:       path + "/" + f.Name(),
			Filetype:       filetype12(fileStat.IsDir()),
			Lastupdatetime: f.ModTime().String(),
			Filename:       f.Name(),
			Filesize:       int(f.Size()),
		}
		fileinfos1 = append(fileinfos1, *fileinffo)
	}
	handleRequests()
}
