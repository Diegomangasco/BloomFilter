# BloomFilter

Welcome to the Bloom Filter Library in Golang! This library provides a robust and efficient implementation of a Bloom Filter, a probabilistic data structure widely used for membership testing within a set. 

## What is a Bloom Filter?

A Bloom Filter is a space-efficient and fast data structure used to determine if an element belongs to a set or not. It's particularly valuable when you need to quickly check membership without the need for expensive data retrieval operations. A Bloom Filter is designed to deliver high performance while maintaining low memory consumption.

## Key Features:

- **Memory Efficiency:** Bloom Filters use minimal memory, making them ideal for applications with large datasets.
  
- **Fast Membership Queries:** They allow for blazing-fast membership queries, enabling quick checks of set membership.

- **False Positives:** While efficient, it's important to note that Bloom Filters may produce false positives, meaning they might mistakenly report that an element is in the set when it's not. However, they never produce false negatives, which makes them useful for various use cases where occasional false positives can be tolerated.

## Use Cases:

Bloom Filters find applications in various domains, including:

- **Database Systems:** They are used to reduce disk I/O operations by determining whether a requested item is likely to be in a cache.

- **Network Protocols:** Bloom Filters are employed for efficient routing table lookups and reducing the load on routers.

- **Spelling Correction:** They help quickly identify whether a word exists in a dictionary, improving spelling correction algorithms.

- **Duplicate Elimination:** In data deduplication and duplicate detection tasks, they help eliminate duplicates efficiently.

## Features
- Create and initialize a Bloom Filter with a specified size and number of hash functions.
- Insert items of type `string` or `int` into the filter.
- Check if an item might be in the filter with the `Contains` method.
- Estimate the false positive rate with the `FalsePositiveRate` method.
- Perform union and intersection operations on Bloom Filters.
- Estimate the number of items present inside the Bloom Filter with the `NumberOfItems` method.

## Installation

To use this Bloom Filter library in your Golang project, follow these steps:

1. Make sure you have Go installed on your system. If not, you can download it from [the official Go website](https://golang.org/dl/).

2. Clone this repository to your local machine or include it as a dependency in your Go project using Go Modules:

   ```shell
   go get github.com/Diegomangasco/BloomFilter
   ```

## Usage

Here's a basic example of how to use the Bloom Filter library in your Go code:

  ```go
  package main

  import (
	  "fmt"
	  "github.com/Diegomangasco/BloomFilter"
  )

  func main() {
	  // Create a new Bloom Filter with a specific capacity and false positive rate
	  filter := bloomfilter.NewBloomFilter(10000, 0.01)

	  // Add items to the filter
	  filter.Add([]byte("item1"))
	  filter.Add([]byte("item2"))

	  // Check if an item may be in the filter (may produce false positives)
	  isInFilter := filter.Contains([]byte("item1"))
	  fmt.Printf("Is 'item1' in the filter? %v\n", isInFilter)

	  isInFilter = filter.Contains([]byte("item3"))
	  fmt.Printf("Is 'item3' in the filter? %v\n", isInFilter)
  }
  ```

## API Reference

For detailed information on how to use the package, please refer to the [API Reference](./APIReference.md).

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.


