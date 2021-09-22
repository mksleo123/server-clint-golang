package solution

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func filetype(c bool) string {
	if c == true {
		return "Dir"
	}
	return "file"
}

func PartA() {
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
		fmt.Println("Systemfilepath: " + path + "/" + f.Name())
		fmt.Println("type: " + filetype(fileStat.IsDir()))
		fmt.Println(f.ModTime().String())
		fmt.Println("filename: " + f.Name())
		fmt.Println(f.Size())
	}
}
