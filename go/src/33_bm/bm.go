package _33_bm

type BMSTRUCT struct {
}

const (
	SIZE = 256 //全局变量
)

//a主串,n主串长度,b模式串,m模式串长度
func BM(a []byte, n int, b []byte, m int) int {
	bc := make([]int, SIZE) //记录模式串中每个字符最后出现的位置
	generateBC(b, m, bc)    //构建坏字符哈希表
	suffix := make([]int, m)
	prefix := make([]bool, m)
	generateGS(b, m, suffix, prefix)
	i := 0 //j表示主串与模式串匹配的第一个字符
	for i <= n-m {
		var j int
		for j = m - 1; j >= 0; j-- {
			if a[i+j] != b[j] {
				break
			}
		}
		if j < 0 { // 匹配成功，返回主串与模式串第一个匹配的字符的位置
			return i
		}
		x := j - bc[int(a[i+j])] //坏字符规则求出的滑动位数
		y := 0
		if j < m-1 { //如果有好后缀
			x = moveByGS(j, m, suffix, prefix) //好后缀规则活动位数
		}
		var max int
		if x > y {
			max = x
		} else {
			max = y
		}
		i = i + max
	}
	return -1
}

// j表示坏字符对应的模式串中的字符下标; m表示模式串长度
func moveByGS(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j //好后缀长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}

	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] == true {
			return r
		}
	}
	return m
}

//将模式串字符下标存入BC散列表
func generateBC(b []byte, m int, bc []int) {
	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}
	for i := 0; i < m; i++ {
		ascii := int(b[i])
		bc[ascii] = i
	}
}

// b表示模式串，m表示长度，suffix，prefix数组事先申请好了,好后缀规则
func generateGS(b []byte, m int, suffix []int, prefix []bool) {
	for i := 0; i < m; i++ { //初始化
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ { // b[0, i]
		j := i
		k := 0                           //公共后缀子串长度
		for j >= 0 && b[j] == b[m-1-k] { // 与b[0, m-1]求公共后缀子串
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 { //如果公共后缀子串也是模式串的前缀子串
			prefix[k] = true
		}
	}
}
