package belajargolangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed cat.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("cat_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}