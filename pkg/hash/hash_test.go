package hash

import (
	"testing"
)

func TestSHA1Hex(t *testing.T) {

	hash := SHA1Hex("https://auth.example.com")
	if len(hash) != 40 {
		t.Errorf("Expected SHA-1 hex to be 40 characters long, got %d", len(hash))
	}

	expect := "cecdaf0304a9dc0fb248cfc8318f489163f974dc"
	if hash != expect {
		t.Errorf("Expected SHA-1 hex to be %s, got %s", expect, hash)
	}

	t.Log("SHA-1 hex:", hash)

}
