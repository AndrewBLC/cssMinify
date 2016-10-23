package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

func main() {
	var filepath string
	fmt.Scan(&filepath)
	cssFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	cssFile, err = m.Bytes("text/css", cssFile)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath, cssFile, 0666)
	if err != nil {
		panic(err)
	}
}
