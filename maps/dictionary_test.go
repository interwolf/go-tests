package maps

import "testing"

func TestDelete(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		// want := "explanation of test"
		var err error

		err = dict.Delete(word)
		assertError(t, err, ErrorNotExistence)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		want := "explanation of test"
		dict := Dictionary{word: want}

		var err error

		err = dict.Delete(word)
		assertError(t, err, nil)

		_, err = dict.Search(word)
		assertError(t, err, ErrorNotFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		want := "explanation of test"
		var err error

		err = dict.Update(word, want)
		assertError(t, err, ErrorNotExistence)

		_, err = dict.Search(word)
		assertError(t, err, ErrorNotFound)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		want := "explanation of test"
		dict := Dictionary{word: "aaa"}

		var err error

		err = dict.Update(word, want)
		assertError(t, err, nil)

		got, _ := dict.Search(word)
		assertStrings(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		want := "explanation of test"
		var err error

		err = dict.Add(word, want)
		assertError(t, err, nil)

		got, _ := dict.Search(word)
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		want := "explanation of test"
		dict := Dictionary{word: want}

		var err error

		err = dict.Add(word, "aaa")
		assertError(t, err, ErrorExistence)

		got, _ := dict.Search(word)
		assertStrings(t, got, want)
	})
}

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "explanation of test"}

	t.Run("existing words", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "explanation of test"

		assertStrings(t, got, want)
	})

	t.Run("not existing words", func(t *testing.T) {
		_, got := dict.Search("testa")

		assertError(t, got, ErrorNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// want should not be nil
func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != nil && got != want {
		t.Errorf("got error %q, want error %q", got, want)
	}
	if got == nil && got != want {
		t.Fatalf("expect to get an error %q but get nil", want)
	}
}
