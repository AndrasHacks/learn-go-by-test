package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Known word", func(t *testing.T) {
		const testKey string = "test"
		const testValue string = "this is just a test"

		dictionary := Dictionary{testKey: testValue}

		got, _ := dictionary.Search(testKey)
		want := testValue
		assertStrings(want, got, t)
	})

	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{"myKnownKey": "My known value"}
		_, err := dictionary.Search("unknown value")

		want := WordNotFoundError
		if err == nil {
			t.Fatal("Expected to get an error!")
		}

		assertErrors(want, err, t)

	})

	t.Run("Add a  non-existent word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("key", "value")

		if err != nil {
			t.Fatal("This must not throw an error!")
		}

		got, _ := dictionary.Search("key")
		want := "value"

		assertStrings(want, got, t)
	})

	t.Run("Add same word twice", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}
		err := dictionary.Add("key", "overwrite!")

		if err == nil {
			t.Fatal("Must throw an error!")
		}

		assertErrors(KeyAlreadyExistsError, err, t)
	})

	t.Run("Update a word", func(t *testing.T) {
		dictionary := Dictionary{"OWL": "Ordinary Wizarding Level"}

		dictionary.Update("OWL", "just an owl!")

		got, _ := dictionary.Search("OWL")

		assertStrings("just an owl!", got, t)
	})

	t.Run("Update non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("Owl", "hu huuhh")

		if err == nil {
			t.Fatal("must throw error")
		}

		assertErrors(WordNotFoundError, err, t)
	})
}

func assertStrings(want string, got string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q, but got %q", want, got)
	}
}

func assertErrors(want error, got error, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q, but got %q", want.Error(), got.Error())
	}
}
