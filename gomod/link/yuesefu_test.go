package link

import "testing"

func TestShowNodeCircle(t *testing.T) {
	ShowNodeCircle()
}

func TestPlayGame(t *testing.T) {
	head := AddNodeCircle(10)
	PlayGame(head, 0, 10)
}
