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

func (s *Slice) remove(index int) *int {
	if index < 0 || index >= len(*s) {
		return nil
	}
	(*s)[3] = 0
	removed := (*s)[index]

	if index > 0 && index < len(*s) {
		s1 := append((*s)[:index-1], (*s)[index+1:]...)
		s = &s1
	} else if index == 0 {
		s1 := make(Slice, 0)
		s = &s1
	} else {
		s1 := (*s)[:index-1]
		s = &s1
	}

	return &removed
}

func main() {
	content, err := readFile("/root/.vimrc")
	if err == nil {
		fmt.Println(content)
	}

	type Animal struct {
		Name string
	}

	dog := new(Animal)
	dog.Name = "dog"
	dog2 := dog
	fmt.Println(dog2 == dog)
	dog2.Name = "dog2"
	fmt.Println(dog.Name)
	fmt.Println(dog2.Name)
	fmt.Println(dog2 == dog)

	fmt.Println("------------------")
	slice := &Slice{1, 2, 3, 4, 5}
	fmt.Println(*slice)
	slice2 := slice
	removed := slice2.remove(0)
	if removed != nil {
		fmt.Println(*removed)
	}
	fmt.Println(*slice)
	fmt.Println(slice == slice2)
}
