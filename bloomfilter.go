package bloomfilter

import (
	"errors"
)

// BloomFilter represents a probabilistic data structure used for membership testing.
type BloomFilter struct {
	array          []uint8 // The array of bits used in the Bloom filter.
	hash_functions uint8   // The number of hash functions used for hashing items.
}

// New creates and initializes a new BloomFilter with the specified length and number of hash functions.
func New(length uint16, hash_functions uint8) (BloomFilter, error) {
	if length <= 0 {
		return BloomFilter{}, errors.New("Negative number for length")
	}

	if hash_functions <= 0 {
		return BloomFilter{}, errors.New("Negative number for hash function")
	}

	actual_length := int(length / 8)

	return BloomFilter{array: make([]uint8, actual_length), hash_functions: hash_functions}, nil
}

// Insert adds an item to the BloomFilter.
// The item parameter can be of any type as it is a generic interface{}.
func Insert(bf *BloomFilter, item interface{}) error {
	// Implement insertion logic here based on the item's type and hash it.
	// Update the bits in the Bloom filter array accordingly.
	return nil
}

// Contains checks if an item is possibly in the BloomFilter.
// The item parameter can be of any type as it is a generic interface{}.
func Contains(bf *BloomFilter, item interface{}) bool {
	// Implement the logic to check if the item is possibly in the Bloom filter.
	// Return true if it might be in the filter, false if it definitely isn't.
}
