package main

import "fmt"

//阶乘
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	res := n * Factorial(n-1)
	return res
}

//台阶
func Taijie(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	res := Taijie(n-1) + Taijie(n-2)
	return res
}

type Fibs struct {
	map1 map[int]int
}

//递归实现斐波那契数列 fib(1) = 1 fib(2) = 1 ,fib(n) = fib(n-1)+fib(n-2)
func (fib *Fibs) Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	res, ok := fib.map1[n]
	if ok {
		return res
	}
	res = fib.Fib(n-1) + fib.Fib(n-2)
	fib.map1[n] = res
	return res
}

func main() {
	fmt.Println(Factorial(4))
	fmt.Println(Taijie(4))
	fib := &Fibs{make(map[int]int, 4)}
	fmt.Println(fib.Fib(4))
}
