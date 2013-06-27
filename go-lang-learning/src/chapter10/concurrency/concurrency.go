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

//	var c chan string = make(chan string)
//	go pinger(c)
//	go ponger(c)
//	go printer(c)

	test()

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

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func test() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()
}
