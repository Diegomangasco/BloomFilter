package bloomfilter

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// murmurHash2 calculates the MurmurHash2 hash of the given byte slice with the provided seed.
func murmurHash2(key []byte, seed uint32) uint32 {
	const (
		m = 0x5bd1e995
		r = 24
	)

	length := len(key)
	h := seed ^ uint32(length)

	for length >= 4 {
		k := binary.LittleEndian.Uint32(key)
		k *= m
		k ^= k >> r
		k *= m

		h *= m
		h ^= k

		key = key[4:]
		length -= 4
	}

	switch length {
	case 3:
		h ^= uint32(key[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(key[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(key[0])
		h *= m
	}

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	return h
}

// BloomFilter represents a probabilistic data structure used for membership testing.
type BloomFilter struct {
	array_size     uint32 // The array's size
	array          []byte // The array of bits used in the Bloom filter.
	hash_functions uint8  // The number of hash functions used for hashing items.
}

// New creates and initializes a new BloomFilter with the specified length and number of hash functions.
func New(length uint16, hash_functions uint8) (*BloomFilter, error) {
	if length <= 0 {
		return nil, errors.New("negative number for length")
	}

	if hash_functions <= 0 {
		return nil, errors.New("negative number for hash function")
	}

	actual_length := int(length/8) + 1
	bf := BloomFilter{array_size: uint32(actual_length), array: make([]byte, actual_length), hash_functions: hash_functions}

	return &bf, nil
}

// GetArray retrieves the array associated with the BloomFilter structure
func (bf *BloomFilter) GetArray() ([]byte, error) {
	if bf.array != nil {
		return bf.array, nil
	}

	return nil, errors.New("structure inizialized badly")
}

// GetHashFunctions retrieves the number of hash functions associated with the BloomFilter structure
func (bf *BloomFilter) GetHashFunctions() (uint8, error) {
	if bf.array != nil {
		return bf.hash_functions, nil
	}

	return 0, errors.New("structure inizialized badly")
}

// Insert adds an item to the BloomFilter.
// The item parameter can be of any type as it is a generic interface{}.
func (bf *BloomFilter) Insert(item interface{}) error {
	// Implement insertion logic here based on the item's type and hash it.
	// Update the bits in the Bloom filter array accordingly.
	if strItem, ok := item.(string); ok {
		fmt.Println("Inserting string:", strItem)
		byteString := []byte(strItem)
		for i := 0; i < int(bf.hash_functions); i++ {
			res := murmurHash2(byteString, uint32(i)) % 32
			fmt.Println(res)
		}
	} else if intItem, ok := item.(int); ok {
		fmt.Println("Inserting int:", intItem)

	} else {
		return errors.New("unsupported item type")
	}
	return nil
}

// Contains checks if an item is possibly in the BloomFilter.
// The item parameter can be of any type as it is a generic interface{}.
func Contains(bf *BloomFilter, item interface{}) bool {
	// Implement the logic to check if the item is possibly in the Bloom filter.
	// Return true if it might be in the filter, false if it definitely isn't.
	return false
}
