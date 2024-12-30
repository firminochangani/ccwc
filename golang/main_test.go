package main

import (
	"bytes"
	"testing"
)

// To be extended with more test cases

func TestCW(t *testing.T) {
	mockFile := bytes.NewReader([]byte("Hello\nworld\nthe world\nis a beautiful place"))
	result, err := getStatsFromReader(mockFile)
	requireNoError(t, err)

	t.Log("total_lines", result.TotalLines)
	t.Log("total_characters", result.TotalCharacters)
	t.Log("total_bytes", result.TotalBytes)
	t.Log("total_words", result.TotalWords)

	requireEqual(t, 4, result.TotalLines)
	requireEqual(t, 8, result.TotalWords)
	requireEqual(t, 42, result.TotalCharacters)
	requireEqual(t, 42, result.TotalBytes)
}

func requireNoError(t *testing.T, err error) {
	if err != nil {
		t.FailNow()
	}
}

func requireEqual[T comparable](t *testing.T, expected T, actual T) {
	if expected != actual {
		t.Logf("expected %v but instead got %v", expected, actual)
		t.FailNow()
	}
}
