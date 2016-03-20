package store_test

import (
	"github.com/avovsya/gokv/store"
	"testing"
)

func TestPut(t *testing.T) {
	var tests = []struct {
		key string
		val string
	}{
		{"key1", "value"},
		{"key1", "anothervalue"},
		{"emptyvalue", ""},
		{"", ""},
		{"", "emptykey"},
		{"space key", "val"},
		{"key", "space value"},
	}

	for _, test := range tests {
		err := store.Put(test.key, test.val)
		if err != nil {
			t.Errorf("PUT(%q, %q) should not return error", test.key, test.val)
		}
	}
}

func TestGetExistingKeys(t *testing.T) {
	var tests = []struct {
		key string
		val string
	}{
		{"key1", "value"},
		{"key1", "anothervalue"},
		{"", "emptykey"},
		{"emptyvalue", ""},
		{"", ""},
		{"space key", "val"},
		{"key", "space value"},
	}

	for _, test := range tests {
		store.Put(test.key, test.val)
		res, err := store.Get(test.key)

		if err != nil {
			t.Errorf("GET(%q) returned error: %v", test.key, err)
		}

		if res != test.val {
			t.Errorf("GET(%q) = %q; expected %q", test.key, res, test.val)
		}
	}
}

func TestGetMissingKeys(t *testing.T) {
	var tests = []string{
		"missing key",
		"another missing key",
		"another_missing_key",
	}

	for _, test := range tests {
		res, err := store.Get(test)

		if err != nil {
			t.Errorf("GET(%q) returned error: %v", test, err)
		}

		if res != "" {
			t.Errorf("GET(%q) = %q; expected empty string", test, res)
		}
	}
}

func TestDeleteExistingKeys(t *testing.T) {
	var tests = []struct {
		key string
		val string
	}{
		{"key1", "value"},
		{"key1", "anothervalue"},
		{"emptyvalue", ""},
		{"", ""},
		{"", "emptykey"},
		{"space key", "val"},
		{"key", "space value"},
	}

	for _, test := range tests {
		store.Put(test.key, test.val)
		err := store.Delete(test.key)
		res, _ := store.Get(test.key)

		if err != nil {
			t.Errorf("Delete(%q) returned error: %v", test.key, err)
		}

		if res != "" {
			t.Errorf("Get(%q) = %q; expected empty value, after Delete operation", test.key, res, test.val)
		}
	}
}

func TestDeleteMissingKeys(t *testing.T) {
	var tests = []string{
		"some missing key",
		"another_42",
	}

	for _, test := range tests {
		err := store.Delete(test)

		if err != nil {
			t.Errorf("Delete(%q) returned error: %v", test, err)
		}
	}
}
