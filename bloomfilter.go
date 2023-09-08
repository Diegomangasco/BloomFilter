package bloomfilter

import (
	"errors"
)

type BloomFilter struct {
	length         int8
	hash_functions []string
}

// Initialize a new BloomFilter structure given the length and the number of hash functions
func new(length int8, hash_functions int8) (BloomFilter, error) {
	if length <= 0 {
		return BloomFilter{}, errors.New("Negative number for length")
	}

	if hash_functions <= 0 {
		return BloomFilter{}, errors.New("Negative number for hash function")
	}

	for i := 0; i < int(hash_functions); i++ {
		// todo
	}

	return BloomFilter{length: length, hash_functions: hash_functions}, nil
}
