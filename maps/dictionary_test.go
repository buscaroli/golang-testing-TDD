package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test"}
	t.Run("tests it returns an existing value", func(t *testing.T) {

		result := "this is a test"
		expect := Search(dictionary, "test")

		assertResult(t, result, expect)
	})

	t.Run("tests it returns a non existing value", func(t *testing.T) {
		result := ""
		expect := Search(dictionary, "nothing")

		assertResult(t, result, expect)
	})
}

func assertResult(t *testing.T, result, expect string) {
	if result != expect {
		t.Errorf("\nresult: %s\nexpect: %s\n", result, expect)
	}
}
