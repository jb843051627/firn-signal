package model

import (
	"crypto/sha256"
	"encoding/hex"
)

func StableID(prefix, value string) string {
	hash := sha256.Sum256([]byte(prefix + ":" + value))
	return prefix + "-" + hex.EncodeToString(hash[:])[:16]
}
