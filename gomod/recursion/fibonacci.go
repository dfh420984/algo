package recursion

type Fibs struct {
	m map[int]int //存储函数值，避免递归函数重复调用
}

func NewFibs(n int) *Fibs {
	return &Fibs{
		m: make(map[int]int),
	}
}

//Fibonacci 递归计算斐波那契数列
func (f *Fibs) Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		f.m[n] = 1
		return f.m[n]
	}
	if v, ok := f.m[n]; ok {
		return v
	}
	f.m[n] = f.Fibonacci(n-1) + f.Fibonacci(n-2)
	return f.m[n]
}
