/**
 * User: guoyao
 * Time: 06-20-2013 12:16
 * Blog: http://www.guoyao.me
 */

package main

import "fmt"

func main() {
//	test()
//	testSlice()
	testMap()
	//	fmt.Println(fibonacci(40))
}

func test() {
	x := [5]float64{1, 2, 3, 4, 5}
	result := float64(1)
	for i := 0; i < len(x); i++ {
		result *= x[i]
	}
	fmt.Println(result)
	fmt.Println(result/float64(len(x)))

	result = 0
	for _, value := range x {
		result += value
	}
	fmt.Println(result)
}

func testSlice() {
	slice1 := []int {1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func testMap() {
	x := make(map[string]int)
	x["key1"] = 1
	fmt.Println(x["1"])
	y := make(map[string]string)
	y["guoyao"] = "Guoyao Wu"
	y["lingxuan"] = "Lingxuan Wu"
	if name, ok := y["guoyao"]; ok {
		fmt.Printf("\"%s\" exsits\n", name)
	}
}

func fibonacci(n uint) (result uint) {
	if n < 2 {
		return n
	}
	return fibonacci(n - 2) + fibonacci(n - 1)
}
