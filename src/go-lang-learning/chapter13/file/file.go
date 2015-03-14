/**
 * User: guoyao
 * Time: 06-27-2013 16:51
 * Blog: http://www.guoyao.me
 */

package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Main() {
//	readFile()
//	createFile()
//	readDirectory()
	walk()
}

func readFile() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func createFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("test")
}

func readDirectory() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func walk() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
