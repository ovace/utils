package fileStats

import (
	"fmt"
	"log"
	"os"
)

func FileInfo(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(fi.Size())
	return err

}
