package recursion

import "testing"

func TestFactorial(t *testing.T) {
	f := Newfactorial()
	res := f.Factorial(5)
	t.Log(res)
	t.Log(f.m)
}
