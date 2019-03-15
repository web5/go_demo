package ranges

import (
	"fmt"
	"testing"
)

func Test_test1(t *testing.T) {
	var i uint64 = 5
	var s = Factorial(i)
	fmt.Printf("%d 的阶乘是 %d\n", i, s)
}
