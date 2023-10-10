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

/*BloomFilter represents a Bloom Filter data structure used for probabilistic membership testing. It contains the following fields:
    - arraySize (Type: uint32): The size of the array in bits.
    - array (Type: []byte): The array of bits used in the Bloom filter.
    - hashFunctions (Type: uint8): The number of hash functions used for hashing items*/
 ```

## Functions

### `NewBloomFilter(capacity int, falsePositiveRate float64) *BloomFilter`

### `GetArray() []byte`

### `GetHashFunctions() uint8`

### `Insert(item interface{}) error`

### `Contains(item interface{}) (bool, error)`

### `FalsePositiveRate() (float32, error)`

### `Union(bf2 *BloomFilter) (*BloomFilter, error)`

### `Intersection(bf2 *BloomFilter) (*BloomFilter, error)`

### `NumberOfItems() (int, error)`