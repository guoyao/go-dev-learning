package main

import "fmt"
import "os"
import "io"
import "bufio"

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
}
