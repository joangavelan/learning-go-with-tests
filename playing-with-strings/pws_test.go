package pws

import (
	"strings"
	"testing"
	"unicode"
)

func TestStrings(t *testing.T) {
	t.Run("test 'Contains' method", func(t *testing.T) {
		got := strings.Contains("seafood", "foo")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 'Count' method", func(t *testing.T) {
		got := strings.Count("seafood", "o")
		want := 2
		assertCorrectNumber(t, got, want)
	})

	t.Run("test 'Cut' method", func(t *testing.T) {
		before, after, found := strings.Cut("Gopher", "ph")
		beforeShouldBe := "Go"
		afterShouldBe :="er"
		foundShouldBe := true

		if before != beforeShouldBe || after != afterShouldBe || found != foundShouldBe {
			t.Errorf("got %q %q %v, want %q %q %v", before, after, found, beforeShouldBe, afterShouldBe, foundShouldBe)
		}
	})

	t.Run("test 'IndexAny' method", func(t *testing.T) {
		got := strings.IndexAny("football", "aeiouy")
		want := 1
		assertCorrectNumber(t, got, want)
	})

	t.Run("test 'IndexByte' method", func(t *testing.T) {
		got := strings.IndexByte("golang", 'g')
		want := 0
		assertCorrectNumber(t, got, want)
	})

	t.Run("test 'LastIndex' method", func(t *testing.T) {
		got := strings.LastIndex("just do it", "do")
		want := 5
		assertCorrectNumber(t, got, want)
	})

	t.Run("test 'LastIndexAny' method", func(t *testing.T) {
		got := strings.LastIndexAny("now it is time", "it")
		want := 11
		assertCorrectNumber(t, got, want)
	})
	
	t.Run("test 'Repeat' method", func(t *testing.T) {
		got := "ba" + strings.Repeat("na", 2)
		want := "banana"
		assertCorrectString(t, got, want)
	})
	
	t.Run("test 'ToLowerSpecial' method", func(t *testing.T) {
		got := strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")
		want := "önnek iş"
		assertCorrectString(t, got, want)
	})

	t.Run("test 'Trim' method", func(t *testing.T) {
		got := strings.Trim("¡Hol@, cómo estás!", "¡!@")
		want := "Hol@, cómo estás"
		assertCorrectString(t, got, want)
	})

	t.Run("test 'TrimFunc' method", func(t *testing.T) {
		got := strings.TrimFunc("¡Hola!", func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		want := "Hola"
		assertCorrectString(t, got, want)
	})

	t.Run("test 'TrimPrefix' method", func(t *testing.T) {
		const formalName = "Mr. Joan"
		const informalName = "Joan"
		const hello = "Hello "

		got := hello + strings.TrimPrefix(formalName, "Mr. ")
		want := hello + informalName
		assertCorrectString(t, got, want)
	})

	t.Run("test 'TrimSpace' method", func(t *testing.T) {
		got := strings.TrimSpace(" Gavelan ")
		want := "Gavelan"
		assertCorrectString(t, got, want)
	})
}

func assertCorrectNumber(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
} 

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
} 