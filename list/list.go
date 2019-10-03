package list

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// list is a representation of the contents of list.json
var list map[string][]string

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

	json.Unmarshal(file, &list)

	return list
}

// WriteList takes an input from the user and writes to a file, list.json
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
