package recursion

import "testing"

func TestFibonacci(t *testing.T) {
	f := NewFibs(5)
	res := f.Fibonacci(5)
	t.Log(res)
	t.Log(f.m)
}
