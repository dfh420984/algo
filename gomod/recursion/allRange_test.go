package recursion

import "testing"

func TestPermute(t *testing.T) {
	p := NewAllRange([]int{1, 2, 3})
	res := p.Permute()
	t.Log(res)
}
