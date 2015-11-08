package main

import (
	"go-lang-learning/util"
	"io"
	"net/http"
	"os"
)

func testGet() {
	res, err := http.Get("http://guoyao.me")
	util.CheckError(err)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func main() {
	testGet()
}
