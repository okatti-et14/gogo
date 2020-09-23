package add

import "testing"

func TestAddNum(t *testing.T) {
	Helphelp(t, 1)
	if AddNum(1, 2) != 3 {
		t.Error("want:3")
	}
}

func Helphelp(t *testing.T, num int) {
	t.Helper()
	if num == 1 {
		t.Fatal("1")
	}
}
