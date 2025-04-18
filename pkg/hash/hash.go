package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1Hex returns the SHA-1 hash of the given input string as a hexadecimal string.
//
// This function is intended for general-purpose string hashing, such as generating
// consistent filenames or cache keys from dynamic input.
//
// Example:
//
//	hash := SHA1Hex("https://example.com")
//	// hash == "d0e2b0..."
func SHA1Hex(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
