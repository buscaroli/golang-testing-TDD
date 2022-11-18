package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("key does not exist")
var ErrKeyExists = errors.New("key already in dictionary")

// Search if a given key exists in the dictionary and returns the value. If the key does not exists it returns an error.
func (d Dictionary) Search(k string) (string, error) {
	result, ok := d[k]

	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

// Adds a key to the dictionary. If the key already exists returns an error.
func (d Dictionary) Add(k, v string) error {

	result, _ := d[k]
	if result != "" {
		return ErrKeyExists
	}

	d[k] = v
	return nil
}
