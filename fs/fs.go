package fs

import (
	"io/ioutil"
	"os"
)

func GetFileContents(filename string) (data []byte, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	data, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	return
}
