package storage

import (
	"encoding/json"
	"io/ioutil"
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

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadList():\ngot:  %v\nwant: %v", got, want)
	}
}

func TestWriteList(t *testing.T) {
	// Normal use case
	WriteList("foo", []string{"double stuf oreos"})

	got := ReadList()
	want := make(map[string][]string, 2)
	want["foo"] = []string{"milk", "double stuf oreos"}
	want["bar"] = []string{"eggs", "bread"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadList():\ngot:  %v\nwant: %v", got, want)
	}

	// User who doesn't have any items yet
	WriteList("baz", []string{"sour skittles", "hershey bars with almonds"})

	got = ReadList()
	want = make(map[string][]string, 2)
	want["foo"] = []string{"milk", "double stuf oreos"}
	want["bar"] = []string{"eggs", "bread"}
	want["baz"] = []string{"sour skittles", "hershey bars with almonds"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadList():\ngot:  %v\nwant: %v", got, want)
	}

	// Write no new items to someone's list
	WriteList("foo", []string{})

	got = ReadList()
	// Exactly the same as the test above
	want = make(map[string][]string, 2)
	want["foo"] = []string{"milk", "double stuf oreos"}
	want["bar"] = []string{"eggs", "bread"}
	want["baz"] = []string{"sour skittles", "hershey bars with almonds"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadList():\ngot:  %v\nwant: %v", got, want)
	}

	// Restore contents of list.json for future tests
	restoreList()
}

func TestClearList(t *testing.T) {
	// It shouldn't be, but ensure that list.json isn't empty
	WriteList("foo", []string{"double stuf oreos"})
	WriteList("baz", []string{"hershey bars with almonds"})

	ClearList()

	got := ReadList()
	want := make(map[string][]string, 0)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadList():\ngot:  %v\nwant: %v", got, want)
	}

	// Restore contents of list.json for future tests
	restoreList()
}

func TestReadPrices(t *testing.T) {
	got := ReadPrices()
	want := make(map[string]float32, 3)
	want["milk"] = float32(2.39)
	want["eggs"] = float32(2.99)
	want["bread"] = float32(3.29)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadPrices():\ngot:  %v\nwant: %v", got, want)
	}

}

func TestWritePrice(t *testing.T) {
	// Normal use case: adding a new price
	WritePrice("double stuf oreos", float32(2.99))

	got := ReadPrices()
	want := make(map[string]float32, 4)
	want["milk"] = float32(2.39)
	want["eggs"] = float32(2.99)
	want["bread"] = float32(3.29)
	want["double stuf oreos"] = float32(2.99)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("WritePrice():\ngot:  %v\nwant: %v", got, want)
	}

	// Overwriting an existing price
	WritePrice("milk", float32(10)) // Chick-fil-a quarantined all cows to please shareholders

	got = ReadPrices()
	want = make(map[string]float32, 4)
	want["milk"] = float32(10)
	want["eggs"] = float32(2.99)
	want["bread"] = float32(3.29)
	want["double stuf oreos"] = float32(2.99)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("WritePrice():\ngot:  %v\nwant: %v", got, want)
	}

	// Restore contents of prices.json for future tests
	restorePrices()
}

// Some tests evaluate data manipulation functionality. Restore the contents
// of list.json where data is stored if a test changed its contents. This way,
// all tests can assume what state list.json is currently in.
func restoreList() {
	originalContents := make(map[string][]string, 2)
	originalContents["foo"] = []string{"milk"}
	originalContents["bar"] = []string{"eggs", "bread"}

	jsonAsBytes, err := json.Marshal(originalContents)
	if err != nil {
		// TODO: Fix this error handling
		panic(err)
	}
	err = ioutil.WriteFile(listPath, jsonAsBytes, 0644)
	if err != nil {
		// TODO: Fix this error handling
		panic(err)
	}
}

// Some tests evaluate data manipulation functionality. Restore the contents
// of prices.json where data is stored if a test changed its contents. This way,
// all tests can assume what state prices.json is currently in.
func restorePrices() {
	originalContents := make(map[string]float32, 3)
	originalContents["milk"] = float32(2.39)
	originalContents["eggs"] = float32(2.99)
	originalContents["bread"] = float32(3.29)

	jsonAsBytes, err := json.Marshal(originalContents)
	if err != nil {
		// TODO: Fix this error handling
		panic(err)
	}
	err = ioutil.WriteFile(pricesPath, jsonAsBytes, 0644)
	if err != nil {
		// TODO: Fix this error handling
		panic(err)
	}
}
