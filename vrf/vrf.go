package vrf

import (
	"fmt"
	"log"
	"math/big"
)

func GetRatio(vrfOutput []byte) {
	t := &big.Int{}
	t.SetBytes(vrfOutput[:])
	fmt.Println("converted to int ", t)

	precision := uint(8 * (len(vrfOutput) + 1))
	max, b, err := big.ParseFloat("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0, precision, big.ToNearestEven)
	if b != 16 || err != nil {
		log.Fatalf("failed to parse big float constant in sortition")
	}

	h := big.Float{}
	h.SetPrec(precision)
	h.SetInt(t)

	ratio := big.Float{}
	cratio, _ := ratio.Quo(&h, max).Float64()
	hval, _ := h.Float64()
	fmt.Println("h ", hval)
	fmt.Println("ratio: ", cratio)
}
