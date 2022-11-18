package maps

type Dictionary map[string]string

type DictionaryError string

const (
	ErrNotFound         = DictionaryError("key does not exist")
	ErrKeyExists        = DictionaryError("key already in dictionary")
	ErrKeyDoesNotExists = DictionaryError("key does not exists")
)

func (e DictionaryError) Error() string {
	return string(e)
}

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

	result := d[k]
	if result != "" {
		return ErrKeyExists
	}

	d[k] = v
	return nil
}

// Updates an existing key. Returns an error if key does not exists.
func (d Dictionary) Update(k, v string) error {
	result := d[k]

	if result == "" {
		return ErrKeyDoesNotExists
	}

	d[k] = v
	return nil
}
