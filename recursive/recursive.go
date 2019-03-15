package ranges

import "fmt"

//阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		fmt.Printf("%d : %d \n", n, result)
		return result
	}
	return 1
}
