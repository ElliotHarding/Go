package main

import (
	"io/ioutil"
)

//File reading
func readFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

//File writing
func fileWrite(path string, toWrite string) error {
	err := ioutil.WriteFile(path, []byte(toWrite), 0644)

	if err != nil {
		return err
	}
	
	return nil
}