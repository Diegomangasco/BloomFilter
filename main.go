package main

import (
	"fmt"

	"github.com/Diegomangasco/BloomFilter/bloomfilter"
)

func main() {
	bf, err := bloomfilter.New(32, 5)
	if err != nil {
		fmt.Println(err)
	}
	bf.Insert(2500)
}
