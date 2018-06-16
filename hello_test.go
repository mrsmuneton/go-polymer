package main

import (
	"testing"
	"time"
)

// Tests
func TestHashBlock(t *testing.T) {
	mock := &Block{Uid: "1", Timestamp: time.Now(), Text: "All legislative Powers herein granted shall be vested in a Congress of the United States", TextValue: "3All legislative Powers herein granted shall be vested in a Congress of the United States", Part: "preamble", PreviousHash: "08045a86ec6dc810e61313b2655bc28069c9f414a273d813a6176ac8e170da8b", Number: "1", Fork: "Original"}

	if len(mock.Hash) > 0 {
		t.Errorf("Hash not empty, equal to %s", mock.Hash)
	}
	hashBlock(mock)
	if len(mock.Hash) != 64 {
		t.Errorf("Unexpected hash, equal to %s", mock.Hash)
	}
}
