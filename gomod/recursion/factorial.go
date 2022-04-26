package recursion

type factorial struct {
	m map[int]int //存储函数值，避免递归函数重复调用
}

func Newfactorial() *factorial {
	return &factorial{
		m: make(map[int]int),
	}
}

//Factorial 递归计算阶乘
func (f *factorial) Factorial(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		f.m[n] = 1
		return f.m[n]
	}
	if v, ok := f.m[n]; ok {
		return v
	}
	f.m[n] = f.Factorial(n-1) * n
	return f.m[n]
}
