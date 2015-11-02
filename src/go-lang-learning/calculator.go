package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	MAX = 50000000000
)

func add(start, end int) uint64 {
	var result uint64 = 0
	for i := start; i <= end; i++ {
		result = result + uint64(i)
	}
	return result
}

func test_add() {
	t0 := time.Now()
	result := add(1, MAX)
	t1 := time.Now()
	fmt.Println("cost time: ", t1.Sub(t0))
	fmt.Println(result)
}

type Segment struct {
	start, end int
}

func toSegments(start, end, segments int) []Segment {
	var result []Segment = make([]Segment, 0)
	var average = (end - start + 1) / segments
	for i := 0; i < segments; i++ {
		var lastEnd = (i+1)*average + start - 1
		if i == segments-1 && lastEnd < end {
			lastEnd = end
		}
		result = append(result, Segment{i*average + start, lastEnd})
	}
	return result
}

func test_toSegments() {
	var segments []Segment = toSegments(2, 14, 3)
	fmt.Println(segments)
}

func add2(start, end int, c chan uint64) {
	c <- add(start, end)
}

func parallelAdd(start, end int) uint64 {
	var cpus int = runtime.NumCPU()
	//if cpus == 1 {
	//	cpus = 100
	//}
	fmt.Println("NumCPU:", cpus)
	var c chan uint64 = make(chan uint64, runtime.NumCPU())
	var segments []Segment = toSegments(start, end, cpus)
	for _, segment := range segments {
		go add2(segment.start, segment.end, c)
	}
	var result uint64 = 0
	for i := 0; i < cpus; i++ {
		result += <-c
	}

	return result
}

func test_parallelAdd() {
	t0 := time.Now()
	result := parallelAdd(1, MAX)
	t1 := time.Now()
	fmt.Println("cost time: ", t1.Sub(t0))
	fmt.Println(result)
}

func main() {
	//test_add()
	test_parallelAdd()
}
