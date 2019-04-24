package utils

import "github.com/willf/bloom"

const (
	PING = 0
	PONG = 1
)

var ppbloom [2]*bloom.BloomFilter
var bloom_count [2]uint
var ppbcurrent uint
var ppbentries uint

func PPbloomInit(n uint, fp float64) {
	ppbentries := n / 2
	err := fp

	ppbloom[PING] = bloom.NewWithEstimates(ppbentries, err)
	ppbloom[PONG] = bloom.NewWithEstimates(ppbentries, err)

	bloom_count[PING] = 0
	bloom_count[PONG] = 0
	
	ppbcurrent = PING
}

func PPbloomCheck(keyAndIV []byte) bool {
	if result := ppbloom[PING].Test(keyAndIV); result == true {
		return result
	}
	if result := ppbloom[PONG].Test(keyAndIV); result == true {
        return result
	} 
	return false
}

func PPbloomAdd(keyAndIV []byte) {
	ppbloom[ppbcurrent].Add(keyAndIV)
	bloom_count[ppbcurrent]++
	if bloom_count[ppbcurrent] >= ppbentries {
		bloom_count[ppbcurrent] = 0
		if ppbcurrent == PING {
			ppbcurrent = PONG
		} else {
			ppbcurrent = PING
		}
		ppbloom[ppbcurrent].ClearAll()
	}
}

func PPbloomClear() {
	ppbloom[PING].ClearAll()
	ppbloom[PONG].ClearAll()
}
