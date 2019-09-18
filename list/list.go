package list

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

// List is a struct representation of the contents of list.json
type List struct {
	Items []string `json:"items"`
}

// ReadList returns the currently-stored grocery list --- a list of strings
func ReadList() []string {
	file, err := ioutil.ReadFile("list.json")
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	list := List{}
	json.Unmarshal(file, &list)

	return list.Items
}
// WriteList takes an input from the user and writes to a file, list.json
func WriteList(items []string) {
	oldList := ReadList()
	oldList = append(oldList, items...)
	newList := List{Items: oldList}
	jsonAsBytes, err := json.Marshal(newList)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}
	err = ioutil.WriteFile("list.json", jsonAsBytes, 0644)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}

}