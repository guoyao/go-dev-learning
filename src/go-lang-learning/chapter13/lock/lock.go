/**
 * User: guoyao
 * Time: 06-28-2013 13:36
 * Blog: http://www.guoyao.me
 */

package lock

import (
	"fmt"
	"sync"
	"time"
)

func Main() {
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}
