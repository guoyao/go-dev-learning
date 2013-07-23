/**
 * User: guoyao
 * Time: 07-23-2013 10:19
 * Blog: http://www.guoyao.me
 */

package closure

import "fmt"

func init() {
	var closures [2]func()
	for i := 0; i < 2; i++ {
		closures[i] = func(ii int) func() {
			return func() {
				fmt.Println(ii)
			}
		}(i)
	}
	closures[0]()
	closures[1]()

	var closures2 [2]func()
	for i := 0; i < 2; i++ {
		val := i
		closures2[i] = func() {
			fmt.Println(val)
		}
	}
	closures2[0]()
	closures2[1]()
}
