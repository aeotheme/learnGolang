package main

import "fmt"

type BS struct {}
func (bs *BS) nilBS() {}
type BF interface { nilBS() }
func nilBS() *BS { return nil }
func nilBF() BF { return nilBS() }
func nilBF2() (bf BF) {
	bf = nilBS()
	return
}
func nilBF3() BF {
	bs := nilBS()
	if bs == nil {
		return nil
	}
	return bs
}

func main() {
	bs := nilBS()
	bf := nilBF()
	bf2 := nilBF2()
	bf3 := nilBF3()
	fmt.Println(bf == nil, bf2 == nil)
	fmt.Printf("%v, %#v\n", bs, bs)
	fmt.Printf("%v, %#v\n", bf, bf)
	fmt.Printf("%v, %#v\n", bf2, bf2)
	fmt.Printf("%v, %#v\n", bf3, bf3)
}
