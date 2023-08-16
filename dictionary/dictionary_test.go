package main

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("dictionary should find word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, err := dict.Search("test")
		want := "this is a test"
		assertNoError(t, err)
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
	t.Run("dictionary should return error when word not found", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, err := dict.Search("unknown")
		want := ""
		assertError(t, ErrNotFound, err)
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
	t.Run("dictionary should support add", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is just a test")

		got, err := dict.Search("test")
		want := "this is just a test"
		assertNoError(t, err)
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
	t.Run("add should not overwrite existing words", func(t *testing.T) {
		dict := Dictionary{"test": "already here"}
		err := dict.Add("test", "this is just a test")
		assertError(t, ErrExistingWord, err)

		got, err := dict.Search("test")
		want := "already here"
		assertNoError(t, err)
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
	t.Run("dictionary should support update", func(t *testing.T) {
		dict := Dictionary{"test": "already here"}
		old, err := dict.Update("test", "this is just a test")
		want := "already here"
		assertNoError(t, err)
		if old != want {
			t.Errorf("got %s, but want %s", old, want)
		}
		want = "this is just a test"
		got, err := dict.Search("test")
		assertNoError(t, err)
		if got != want {
			t.Errorf("got %s, but want %s", old, want)
		}
	})
	t.Run("update should return error if word is not found", func(t *testing.T) {
		dict := Dictionary{"test": "already here"}
		old, err := dict.Update("another", "this is just a test")
		want := ""
		assertError(t, ErrWordNotFound, err)
		if old != want {
			t.Errorf("got %s, but want %s", old, want)
		}
	})
	t.Run("dictionary should support delete", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		old, err := dict.Delete("test")
		want := "this is just a test"
		assertNoError(t, err)
		if old != want {
			t.Errorf("got %s, but want %s", old, want)
		}
	})
	t.Run("delete should return error if word is not found", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		_, err := dict.Delete("unknown")
		assertError(t, ErrWordNotFound, err)
	})
}

func assertError(t *testing.T, want error, got error) {
	t.Helper()
	if got == nil {
		t.Errorf("got is nil")
	}
	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got is not nil: %s", got)
	}
}
