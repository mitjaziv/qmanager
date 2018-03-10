// Package UUID provides unique id functionality.
//
// Disclaimer: I would normally use some kind of UUID library for example: https://github.com/satori/go.uuid
// but as there was request not to use external libraries I implemented, primitive unique id functionality
// https://en.wikipedia.org/wiki/Universally_unique_identifier
// https://tools.ietf.org/html/rfc4122.html
package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

type UUID string

// New generates a random UUID according to RFC 4122
func NewUUID() UUID {
	uuid := make([]byte, 16)

	// Read random bytes from crypto/rand into uuid
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic(err)
	}

	// Set variant bits for RFC 4122 version

	// Clear first 2 bits.
	uuid[8] = uuid[8] &^ 0xc0

	// Set most significant bit to 1
	uuid[8] = uuid[8] | 0x80

	// Set UUID Version to 4 (pseudo-random)
	// We need to set version bit to 0x40

	// Clear first 4 bits.
	uuid[6] = uuid[6] &^ 0xf0

	// Set bits to  0 1 0 0 for version 4.
	uuid[6] = uuid[6] | 0x40

	return UUID(fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]))
}
