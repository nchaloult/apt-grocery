package list

import "testing"

func TestReadList(t *testing.T) {
	got := ReadList()
	want := []string{"testing", "testing2"}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("ReadList(): got: %q, want: %q", got[i], want[i])
		}
	}
}
