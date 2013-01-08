package main

import (
	"fabriciojs/murmurhash3"
	"fmt"
	"time"
	"flag"
)

var test = flag.Int("test", 0, "test")

const K = 21474836 * 10

func main() {
	flag.Parse()

	fmt.Println("murmurhash")

	hashes := make([][2]uint64, *test)

	start := time.Now()
	for i := 0; i < *test; i++ {
		hashes[i] = murmurhash3.Murmur3F([]byte(fmt.Sprintf("key_%d", i)), 1287253131031037)
	}

	fmt.Println("generated in: ", time.Since(start))

	start = time.Now()

	fmt.Println("analysis colisions...")

	colisions0 :=  make(map[uint64]int)
	colisions1 :=  make(map[uint64]int)
	
	for i := 0; i < *test; i++ {
		hash := murmurhash3.Murmur3F([]byte(fmt.Sprintf("key_%d", i)), 1287253131031037)

		colisions0[hash[0] % K]++
		colisions1[hash[1] % K]++
	}

	c0, c1 := 0, 0

	for _, c := range colisions0 {
		if c > 1 {
			c0 += c - 1
		}
	}

	for _, c := range colisions1 {
		if c > 1 {
			c1 += c - 1
		}
	}

	fmt.Println("analysed in: ", time.Since(start))

	fmt.Println("colisions")
	fmt.Println("hash[0]", c0)
	fmt.Println("hash[1]", c1)
}
