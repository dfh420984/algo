package recursion

type AllRange struct {
	nums   []int   //需要全排列数组 []int{1,2,3}
	path   []int   //路径，深度优先遍历存储路径
	result [][]int //最终全排列集合结果
	used   []bool  //递归时判断集合中是否存在重复元素
}

func NewAllRange(nums []int) *AllRange {
	return &AllRange{
		nums: nums,
		path: make([]int, 0),
		used: make([]bool, 3),
	}
}

//Trackingback 递归函数回溯算法
func (a *AllRange) Trackingback(nums []int) {
	//递归终止条件,叶子节点中的元素个数为nums的长度
	if len(a.path) == len(nums) {
		a.result = append(a.result, append([]int{}, a.path...))
		return
	}
	//横向遍历
	for i := 0; i < len(nums); i++ {
		if a.used[i] { //如果path中存在该元素，跳过
			continue
		}
		a.used[i] = true
		a.path = append(a.path, nums[i])
		//回溯递归调用，开始深度优先遍历
		a.Trackingback(nums)
		//恢复状态
		a.path = a.path[:len(a.path)-1]
		a.used[i] = false
	}
}

//Permute 全排列
func (a *AllRange) Permute() [][]int {
	a.Trackingback(a.nums)
	return a.result
}
