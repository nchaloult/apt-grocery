package list

import (
	"reflect"
	"testing"
)

func TestReadList(t *testing.T) {
	got := ReadList()

	want := make(map[string][]string, 2)
	want["foo"] = []string{"milk"}
	want["bar"] = []string{"eggs", "bread"}
	want["baz"] = []string{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadList():\ngot:  %v\nwant: %v", got, want)
	}
}
