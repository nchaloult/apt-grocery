package list

import (
	"encoding/json"
	"io/ioutil"
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
