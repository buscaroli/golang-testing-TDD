package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("tests it returns an existing value", func(t *testing.T) {

		result, _ := dictionary.Search("test")
		expect := "this is a test"

		assertResult(t, result, expect)
	})

	t.Run("tests it returns a non existing value", func(t *testing.T) {
		_, errMsg := dictionary.Search("nothing")

		assertError(t, errMsg, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("tests it can add a new key", func(t *testing.T) {
		dic_k := "one"
		dic_v := "first"

		dictionary.Add(dic_k, dic_v)

		result, err := dictionary.Search(dic_k)
		expect := dic_v

		assertResult(t, result, expect)
		assertError(t, err, nil)
	})

	t.Run("tests it returns an error if key already exists", func(t *testing.T) {
		errMsg := dictionary.Add("test", "already there")

		// Checks error is returned.
		assertError(t, errMsg, ErrKeyExists)

		// Checks original key has not been updated.
		assertResult(t, dictionary["test"], "this is a test")
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("tests an existing key can be updated", func(t *testing.T) {
		dic_k := "test"
		dic_v := "updated key"
		err := dictionary.Update(dic_k, dic_v)

		assertResult(t, dictionary[dic_k], dic_v)
		assertError(t, err, nil)
	})

	t.Run("tests trying to update a non existent key will return an error", func(t *testing.T) {
		dic_k := "ten"
		dic_v := "tenth"
		err := dictionary.Update(dic_k, dic_v)

		assertError(t, err, ErrKeyDoesNotExists)
		assertResult(t, dictionary[dic_k], "")
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test", "two": "second"}

	t.Run("tests deleting an existing key", func(t *testing.T) {
		dic_k := "two"
		err := dictionary.Delete(dic_k)

		assertError(t, err, nil)
		assertResult(t, dictionary[dic_k], "")
	})

	t.Run("tests trying to delete a non existent key will return an error", func(t *testing.T) {
		dic_k := "three"
		err := dictionary.Delete(dic_k)

		assertError(t, err, ErrKeyDoesNotExists)
		assertResult(t, dictionary[dic_k], "")
	})
}

func assertResult(t testing.TB, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("\nresult: %s\nexpect: %s\n", result, expect)
	}
}

func assertError(t testing.TB, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("\nresult: %s\nexpect: %s\n", result, expect)
	}
}
