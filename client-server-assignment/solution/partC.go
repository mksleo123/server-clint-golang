package solution

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"../types1"
)

var fileinfos []types1.FileInfo

func filetype1(c bool) string {
	if c == true {
		return "Dir"
	}
	return "file"
}

func PartC() {
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
			Filetype:       filetype1(fileStat.IsDir()),
			Lastupdatetime: f.ModTime().String(),
			Filename:       f.Name(),
			Filesize:       int(f.Size()),
		}
		fileinfos = append(fileinfos, *fileinffo)

	}

	b, err := json.Marshal(fileinfos)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
		return
	}

	fmt.Println(string(b))
}
