package list

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var listPath string

func init() {
	runEnv := os.Getenv("RUN_ENV")

	if runEnv == "prod" {
		listPath = "list.json"
	} else {
		listPath = "list/list.json"
	}
}

// ReadList returns the currently-stored grocery list --- a list of strings
func ReadList() map[string][]string {
	file, err := ioutil.ReadFile(listPath)
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	// list is a representation of the contents of list.json
	var list map[string][]string
	json.Unmarshal(file, &list)

	return list
}

// WriteList makes an amendment to a file, taking an input from the user and
// writing it to a file: list.json
func WriteList(user string, items []string) {
	list := ReadList()
	list[user] = append(list[user], items...)
	jsonAsBytes, err := json.Marshal(list)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}
	err = ioutil.WriteFile(listPath, jsonAsBytes, 0644)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}
}

// ClearList wipes the contents of a file: list.json.
//
// Can't call WriteList here, because WriteList appends the provided new items
// with the existing contents of the list.
func ClearList() {
	jsonAsBytes, err := json.Marshal(make(map[string]string, 0))
	if err != nil {
		// TODO: Fix this error handling
		panic(err)
	}
	err = ioutil.WriteFile(listPath, jsonAsBytes, 0644)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}
}
