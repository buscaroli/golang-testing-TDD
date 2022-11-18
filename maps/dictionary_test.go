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
		_, result := dictionary.Search("nothing")

		assertError(t, result, ErrNotFound)
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
