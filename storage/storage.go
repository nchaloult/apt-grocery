package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Relative path at which list.json persistent storage file can be found
var listPath string
// Relative path at which prices.json persistent storage file can be found
var pricesPath string

// Initialize paths to .json files we're using for persistent storage.
// Depends on what environment we're running in (dev or prod)
func init() {
	runEnv := os.Getenv("RUN_ENV")

	if runEnv == "prod" {
		listPath = "list.json"
		pricesPath = "prices.json"
	} else {
		listPath = "storage/list.json"
		pricesPath = "storage/prices.json"
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

// ReadPrices returns the currently-stored prices for commonly-purchased items
func ReadPrices() map[string]float32 {
	file, err := ioutil.ReadFile(pricesPath)
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	// prices is a representation of the contents of prices.json
	var prices map[string]float32
	json.Unmarshal(file, &prices)

	return prices
}

// WritePrice adds a new price for a commonly-purchased item in prices.json.
// If the item involved is already present in prices.json, then that price is
// overwritten.
func WritePrice(item string, price float32) {
	prices := ReadPrices()
	prices[item] = price
	jsonAsBytes, err := json.Marshal(prices)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}
	err = ioutil.WriteFile(pricesPath, jsonAsBytes, 0644)
	if err != nil {
		//TODO: Fix this error handling
		panic(err)
	}
}
