package bloomfilter

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

const SIZEOFUINT8 = 8

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
	arraySize     uint32 // The array's size
	array         []byte // The array of bits used in the Bloom filter.
	hashFunctions uint8  // The number of hash functions used for hashing items.
}

// New creates and initializes a new BloomFilter with the specified length and number of hash functions.
func NewBloomFilter(length uint16, hashFunctions uint8) (*BloomFilter, error) {
	if length <= 0 {
		return nil, errors.New("negative number for length")
	}

	if hashFunctions <= 0 {
		return nil, errors.New("negative number for hash function")
	}

	actual_length := int(length/8) + 1
	bf := BloomFilter{arraySize: uint32(length), array: make([]byte, actual_length), hashFunctions: hashFunctions}

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
		return bf.hashFunctions, nil
	}

	return 0, errors.New("structure inizialized badly")
}

// Insert adds an item to the BloomFilter.
// The item parameter can be of any type as it is a generic interface{}.
func (bf *BloomFilter) Insert(item interface{}) error {
	// Implement insertion logic here based on the item's type and hash it.
	// Update the bits in the Bloom filter array accordingly.

	var byteArray []byte
	if strItem, ok := item.(string); ok {
		fmt.Println("Inserting string:", strItem)
		byteArray = []byte(strItem)
	} else if intItem, ok := item.(int); ok {
		fmt.Println("Inserting int:", intItem)
		byteArray = make([]byte, 4)
		binary.LittleEndian.PutUint32(byteArray, uint32(intItem))
	} else {
		return errors.New("unsupported item type")
	}

	for i := 0; i < int(bf.hashFunctions); i++ {
		res := murmurHash2(byteArray, uint32(i)) % bf.arraySize // Compute the murmur hashing
		arrayPos := int(res / SIZEOFUINT8)                      // Array's cell
		cellPos := int(res) - arrayPos*SIZEOFUINT8 - 1          // Bit inside the specific cell
		if cellPos == -1 {
			cellPos = 7
		}
		bf.array[arrayPos] |= 1 << uint(cellPos) // Set the bit to 1
	}

	return nil
}

// Contains checks if an item is possibly in the BloomFilter.
// The item parameter can be of any type as it is a generic interface{}.
// Return true if it might be in the filter, false if it definitely isn't.
func (bf *BloomFilter) Contains(item interface{}) (bool, error) {
	// Implement the logic to check if the item is possibly in the Bloom filter.

	var byteArray []byte
	if strItem, ok := item.(string); ok {
		fmt.Println("Checking string:", strItem)
		byteArray = []byte(strItem)
	} else if intItem, ok := item.(int); ok {
		fmt.Println("Checking int:", intItem)
		byteArray = make([]byte, 4)
		binary.LittleEndian.PutUint32(byteArray, uint32(intItem))
	} else {
		return false, errors.New("unsupported item type")
	}

	for i := 0; i < int(bf.hashFunctions); i++ {
		res := murmurHash2(byteArray, uint32(i)) % bf.arraySize // Compute the murmur hashing
		arrayPos := int(res / SIZEOFUINT8)                      // Array's cell
		cellPos := int(res) - arrayPos*SIZEOFUINT8 - 1          // Bit inside the specific cell
		if cellPos == -1 {
			cellPos = 7
		}
		bitMask := byte(1 << uint(cellPos))
		value := bf.array[arrayPos] // Get the value of the specific cell in the BloomFilter's array
		value &= bitMask
		if value == 0 {
			return false, nil
		}
	}

	return true, nil
}

// FalsePositiveRate returns the false positive rate based on the number of hash functions and the filter size of the BloomFilter.
func (bf *BloomFilter) FalsePositiveRate() (float64, error) {
	if bf.array == nil {
		return -1.0, errors.New("structure inizialized badly")
	}

	elements, err := bf.NumberOfItems()
	if err != nil {
		return -1.0, err
	}
	fpr := math.Pow(1-math.Exp(float64(elements)*float64(bf.hashFunctions)/(float64(bf.arraySize))), float64(bf.hashFunctions))

	return fpr, nil
}

// Union creates a new Bloom Filter representing the union of two Bloom Filters.
func (bf1 *BloomFilter) Union(bf2 *BloomFilter) ([]byte, error) {

	if bf1.array == nil || bf2.array == nil {
		return make([]byte, 0), errors.New("structure inizialized badly")
	}

	var arr []byte
	var major []byte
	var minor []byte
	var major_size int
	var minor_size int
	size1 := int(bf1.arraySize/8) + 1
	size2 := int(bf2.arraySize/8) + 1
	if size1 > size2 {
		major = bf1.array
		minor = bf2.array
		major_size = size1
		minor_size = size2
	} else {
		major = bf2.array
		minor = bf1.array
		major_size = size2
		minor_size = size1
	}

	for i := 0; i < major_size; i++ {
		if i < minor_size {
			arr[i] = major[i] | minor[i]
		} else {
			arr[i] = major[i]
		}
	}

	return arr, nil
}

// Intersection creates a new Bloom Filter representing the intersection of two Bloom Filters.
func (bf1 *BloomFilter) Intersection(bf2 *BloomFilter) ([]byte, error) {

	if bf1.array == nil || bf2.array == nil {
		return make([]byte, 0), errors.New("structure inizialized badly")
	}

	var arr []byte
	var major []byte
	var minor []byte
	var major_size int
	var minor_size int
	size1 := int(bf1.arraySize/8) + 1
	size2 := int(bf2.arraySize/8) + 1
	if size1 > size2 {
		major = bf1.array
		minor = bf2.array
		major_size = size1
		minor_size = size2
	} else {
		major = bf2.array
		minor = bf1.array
		major_size = size2
		minor_size = size1
	}

	for i := 0; i < major_size; i++ {
		if i < minor_size {
			arr[i] = major[i] & minor[i]
		} else {
			arr[i] = 0
		}
	}

	return arr, nil
}

// NumberOfItems estimates the number of items present inside the bloom filter
func (bf *BloomFilter) NumberOfItems() (int, error) {
	if bf.array != nil {
		m := int(bf.arraySize)
		k := int(bf.hashFunctions)
		x := 0
		for i := 0; i < int(bf.arraySize/8)+1; i++ {
			item := bf.array[i]
			count := 0
			for item > 0 {
				count += int(item & 1)
				item >>= 1
			}
			x += count
		}
		a := -float64(m) / float64(k)
		b := math.Log(float64(1) - float64(x)/float64(m))
		return int(a * b), nil
	}

	return 0, errors.New("structure inizialized badly")
}
