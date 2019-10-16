package storage

import (
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"
)

// During tests, sets file path to the project root directory.
// Otherwise, running `go test ./...` in the project root fails because
// list.json can't be found.
//
// https://brandur.org/fragments/testing-go-project-root
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

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
