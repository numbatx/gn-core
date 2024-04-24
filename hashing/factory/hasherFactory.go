package factory

import (
	"github.com/numbatx/gn-core/hashing"
	"github.com/numbatx/gn-core/hashing/blake2b"
	"github.com/numbatx/gn-core/hashing/keccak"
	"github.com/numbatx/gn-core/hashing/sha256"
)

// NewHasher will return a new instance of hasher based on the value stored in config
func NewHasher(name string) (hashing.Hasher, error) {
	switch name {
	case "sha256":
		return sha256.NewSha256(), nil
	case "keccak":
		return keccak.NewKeccak(), nil
	case "blake2b":
		return blake2b.NewBlake2b(), nil
	}

	return nil, ErrNoHasherInConfig
}
