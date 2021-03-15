package file

import (
	"fmt"
	"testing"
)

func TestFileExist(t *testing.T) {
	path := "./a.txt"
	f, err := Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var b []byte
	n, err := f.Read(b)
	if err != nil {
		fmt.Println(err, n)
	}
	defer f.Close()
}

func TestFindNotExistKeys(t *testing.T) {
	path := "src.log"
	keys := []string{
		"d",
		"hello",
	}
	output, err := FindNotExistKeys(path, keys)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("h1")
	fmt.Println(output)
	fmt.Println("h1")
}
