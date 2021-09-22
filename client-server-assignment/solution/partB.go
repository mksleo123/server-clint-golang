package solution

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"../types1"
)

func filetype1s(c bool) string {
	if c == true {
		return "Dir"
	}
	return "file"
}

func PartB() {
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
			Filetype:       filetype1s(fileStat.IsDir()),
			Lastupdatetime: f.ModTime().String(),
			Filename:       f.Name(),
			Filesize:       int(f.Size()),
		}

		b, err := json.MarshalIndent(fileinffo, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

	}
}
