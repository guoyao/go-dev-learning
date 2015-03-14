/**
 * User: guoyao
 * Time: 06-28-2013 13:29
 * Blog: http://www.guoyao.me
 */

package arguments

import (
	"fmt"
	"flag"
	"math/rand"
)

func Main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
