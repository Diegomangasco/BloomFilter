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
	bf.Insert(5122)
	bf.Insert(10)
	bf.Insert(161245)
	bf.Insert(80000)
	array, err := bf.GetArray()
	if err == nil {
		fmt.Println(array)
	}

	res, err := bf.Contains(20)
	if err == nil {
		fmt.Println(res)
	}

}
