package main

import "fmt"
import "os"
import "io"
import "bufio"
import "crypto/hmac"
import "crypto/sha256"

func readFile(path string) (result string, err error) {
	result = ""
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		fmt.Println("Failed to open file ", path)
		return
	}

	br := bufio.NewReader(file)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		result += string(line) + "\n"
	}

	return
}

type Slice []int

func (s Slice) remove(index int) (*int, Slice) {
	if index < 0 || index >= len(s) {
		return nil, s
	}
	removed := s[index]

	s = append(s[:index], s[index+1:]...)

	return &removed, s
}

func toHmac(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func hello(c chan interface{}) {
	c <- "1"
	fmt.Println("hello")
}

type Person struct {
	Age   int
	Names []string
}

func main() {
	//content, err := readFile(os.Getenv("HOME") + "/.vimrc")
	//if err == nil {
	//	fmt.Println(content)
	//}

	c := make(chan interface{}, 1)
	go hello(c)
	i := <-c
	fmt.Println(i)

	person := &Person{
		20, []string{"hell\"dsdf'o", "world"},
	}

	fmt.Println(*person)

	fmt.Println(string(toHmac([]byte("bad522c2126a4618a8125f4b6cf6356f"), []byte("bce-auth-v1/0b0f67dfb88244b289b72b142befad0c/2015-04-27T08:23:49Z/1800"))))
}
