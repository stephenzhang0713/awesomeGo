package main

import (
	"fmt"
	"github.com/bytedance/gopkg/lang/fastrand"
)

func main() {
	a := fastrand.Int()
	fmt.Printf("rand int: %d\n", a)

	b := fastrand.Intn(100)
	fmt.Printf("rand intn: %d\n", b)

	c := fastrand.Uint32()
	fmt.Printf("rand uint32: %d\n", c)

	d := fastrand.Uint64()
	fmt.Printf("rand uint64: %d\n", d)

	e := fastrand.Float32()
	fmt.Printf("rand float32: %f\n", e)

	f := fastrand.Float64()
	fmt.Printf("rand float64: %f\n", f)

	g := fastrand.Perm(10)
	fmt.Printf("rand perm: %v\n", g)

	read, err := fastrand.Read([]byte("hello"))
	if err != nil {
		return
	}
	fmt.Printf("rand read: %v\n", read)

}
