package main

import (
	"github.com/willf/bloom"
	"github.com/waysaku/bloomfilter/origin"
	"fmt"
)

func main()  {
	fmt.Println("START")
	origin.Filter()

	n := uint(1000)
	filter := bloom.New(20 * n, 5) // load of 20, 5 keys
	filter.Add([]byte("Love"))

	if filter.Test([]byte("Love")) {
		fmt.Println("DETECT!!")
	}
	if filter.Test([]byte("HOGE")) {
		fmt.Println("DETECT!!")
	}
}


