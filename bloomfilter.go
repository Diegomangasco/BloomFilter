package bloomfilter

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

const SIZEOFUINT8 = 8

func murmurHash3(key []byte, seed uint32) uint32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)

	hash := seed
	data := key
	length := len(key)
	nblocks := length / 4

	for i := 0; i < nblocks; i++ {
		k := *(*uint32)(unsafe.Pointer(&data[i*4]))
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2

		hash ^= k
		hash = (hash << r2) | (hash >> (32 - r2))
		hash = hash*m + n
	}

	tail := data[nblocks*4:]
	k1 := uint32(0)

	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = (k1 << r1) | (k1 >> (32 - r1))
		k1 *= c2
		hash ^= k1
	}

	hash ^= uint32(length)
	hash ^= (hash >> 16)
	hash *= 0x85ebca6b
	hash ^= (hash >> 13)
	hash *= 0xc2b2ae35
	hash ^= (hash >> 16)

	return hash
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
		res := murmurHash3(byteArray, uint32(i))
		res %= bf.arraySize
		arrayPos := int(res / SIZEOFUINT8)             // Array's cell
		cellPos := int(res) - arrayPos*SIZEOFUINT8 - 1 // Bit inside the specific cell
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
		res := murmurHash3(byteArray, uint32(i))
		res %= bf.arraySize
		arrayPos := int(res / SIZEOFUINT8)             // Array's cell
		cellPos := int(res) - arrayPos*SIZEOFUINT8 - 1 // Bit inside the specific cell
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
	exp := math.Exp(float64(elements) * float64(bf.hashFunctions) / (float64(bf.arraySize)))
	diff := 1 - exp
	if diff < 0 {
		diff = 0
	}
	fpr := math.Pow(diff, float64(bf.hashFunctions))

	return fpr, nil
}

// Union creates a new Bloom Filter representing the union of two Bloom Filters.
func (bf1 *BloomFilter) Union(bf2 *BloomFilter) (*BloomFilter, error) {

	if bf1.array == nil || bf2.array == nil {
		return &BloomFilter{}, errors.New("structure inizialized badly")
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
			arr = append(arr, major[i]|minor[i])
		} else {
			arr = append(arr, major[i])
		}
	}

	bf_union, err := NewBloomFilter(uint16(major_size-1)*SIZEOFUINT8, bf1.hashFunctions)
	if err != nil {
		return &BloomFilter{}, errors.New("structure inizialized badly")
	}

	bf_union.array = arr

	return bf_union, nil
}

// Intersection creates a new Bloom Filter representing the intersection of two Bloom Filters.
func (bf1 *BloomFilter) Intersection(bf2 *BloomFilter) (*BloomFilter, error) {

	if bf1.array == nil || bf2.array == nil {
		return &BloomFilter{}, errors.New("structure inizialized badly")
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
			arr = append(arr, major[i]&minor[i])
		} else {
			arr = append(arr, 0)
		}
	}

	bf_intersection, err := NewBloomFilter(uint16(major_size-1)*SIZEOFUINT8, bf1.hashFunctions)
	if err != nil {
		return &BloomFilter{}, errors.New("structure inizialized badly")
	}

	bf_intersection.array = arr

	return bf_intersection, nil
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
		return int(math.Ceil(a * b)), nil
	}

	return 0, errors.New("structure inizialized badly")
}
