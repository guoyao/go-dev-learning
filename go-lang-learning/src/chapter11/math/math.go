/**
 * User: guoyao
 * Time: 06-27-2013 13:37
 * Blog: http://www.guoyao.me
 */

package math

// Finds the average of a series of numbers
func Average(value []float64) float64 {
	var total float64
	for _, x := range value {
		total += x
	}
	return total / float64(len(value))
}
