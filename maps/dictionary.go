package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("key does not exist")

func (d Dictionary) Search(k string) (string, error) {
	result, ok := d[k]

	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}
