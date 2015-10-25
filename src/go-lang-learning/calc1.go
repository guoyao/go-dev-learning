package main

import "fmt"
import _ "time"
import "reflect"

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}

func main() {
	/*
		t1 := time.Now()
		count := 0
		for i := 0; i < 9000000000; i++ {
			count += i
		}
		t2 := time.Now()
		fmt.Printf("cost:%d, count:%d\n", t2.Sub(t1), count)
	*/

	// arr := []int{1, 2, 3}

	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	var nu int = 10
	fmt.Println(string(nu))
}
