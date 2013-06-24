/**
 * User: guoyao
 * Time: 06-24-2013 17:07
 * Blog: http://www.guoyao.me
 */

package concurrency

import (
	"fmt"
	"time"
	"math/rand"
)

func Run() {
//	for i := 0; i < 10; i++ {
//		go f(i)
//	}

	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
