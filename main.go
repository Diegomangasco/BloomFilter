package main

import (
	"fmt"

	"github.com/Diegomangasco/BloomFilter/bloomfilter"
)

func main() {
	bf, err := bloomfilter.NewBloomFilter(32, 5)
	if err != nil {
		fmt.Println(err)
	}
	bf.Insert(5122)
	bf.Insert(10)
	bf.Insert(161245)
	bf.Insert(80000)
	n, err := bf.NumberOfItems()
	if err == nil {
		fmt.Println(n)
	}

}
