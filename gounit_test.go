package main

import "testing"

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected Error, %s", err)
	}
}
func AssertEqualStrings(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Strings don't match;\n got      '%s'\n expected '%s'", actual, expected)
	}
}
func AssertEqualInts(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Integers don't match;\n got      '%d'\n expected '%d'", actual, expected)
	}
}
