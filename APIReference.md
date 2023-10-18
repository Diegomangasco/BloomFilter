# Bloom Filter Package API Reference

This document provides detailed information about the API of the `bloomfilter` package for Go (Golang). The package implements a Bloom Filter data structure for probabilistic membership testing.

## Table of Contents

1. [Types](#types)
   - [BloomFilter](#bloomfilter)
2. [Functions](#functions)
   - [`NewBloomFilter(capacity int, falsePositiveRate float64) *BloomFilter`](#newbloomfiltercapacity-int-falsepositiverate-float64-bloomfilter)
   - [`GetArray() []byte`](#getarray-byte)
   - [`GetHashFunctions() uint8`](#gethashfunctions-uint8)
   - [`Insert(item interface{}) error`](#insertitem-interface-error)
   - [`Contains(item interface{}) (bool, error)`](#containsitem-interface-bool-error)
   - [`FalsePositiveRate() (float32, error)`](#falsepositiverate-float32-error)
   - [`Union(bf2 *BloomFilter) (*BloomFilter, error)`](#unionbf2-bloomfilter-bloomfilter-error)
   - [`Intersection(bf2 *BloomFilter) (*BloomFilter, error)`](#intersectionbf2-bloomfilter-bloomfilter-error)
   - [`NumberOfItems() (int, error)`](#numberofitems-int-error)

---

## Types

### `BloomFilter`

```go
type BloomFilter struct {
    arraySize     uint32
    array         []byte
    hashFunctions uint8
}
 ```
BloomFilter represents a Bloom Filter data structure used for probabilistic membership testing. It contains the following fields:

   - arraySize (Type: uint32): The size of the array in bits
   - array (Type: []byte): The array of bits used in the Bloom filter
   - hashFunctions (Type: uint8): The number of hash functions used for hashing items


## Functions

### `NewBloomFilter(capacity int, hashFunctions uint8) *BloomFilter`

Creates and initializes a new BloomFilter with the specified length and number of hash functions.

- `length` (uint16): The desired length of the Bloom Filter
- `hashFunctions` (uint8): The number of hash functions to use

Returns a pointer to the newly created BloomFilter and an error if the initialization fails

### `GetArray() []byte`

Retrieves the array associated with the BloomFilter structure.

Returns the array and an error if the structure is initialized improperly.

### `GetHashFunctions() uint8`

Retrieves the number of hash functions associated with the BloomFilter structure.

Returns the number of hash functions and an error if the structure is initialized improperly.

### `Insert(item interface{}) error`

Adds an item to the BloomFilter. The item parameter can be of any type as it is a generic `interface{}`.

### `Contains(item interface{}) (bool, error)`

Checks if an item is possibly in the BloomFilter. The item parameter can be of any type as it is a generic interface{}.

Returns `true` if the item might be in the filter, `false` if it definitely isn't, and an error if the item type is not supported.

### `FalsePositiveRate() (float32, error)`

Calculates and returns the false positive rate based on the number of hash functions and the filter size of the BloomFilter.

Returns the false positive rate and an error if the structure is initialized improperly.

### `Union(bf2 *BloomFilter) (*BloomFilter, error)`

Creates a new Bloom Filter representing the union of two Bloom Filters.

Returns a new Bloom Filter representing the union of the two input Bloom Filters and an error if the structure is initialized improperly.

### `Intersection(bf2 *BloomFilter) (*BloomFilter, error)`

Creates a new Bloom Filter representing the intersection of two Bloom Filters.

Returns a new Bloom Filter representing the intersection of the two input Bloom Filters and an error if the structure is initialized improperly.

### `NumberOfItems() (int, error)`

Estimates the number of items present inside the Bloom Filter.

Returns the estimated number of items and an error if the structure is initialized improperly.