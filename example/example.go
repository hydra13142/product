package main

import "github.com/hydra13142/product"

func main() {
	x := product.NewArrange(4, 5)
	for !x.Over() {
		for i := 0; i < x.Length(); i++ {
			print(x.Pick(i), " ")
		}
		println("")
		x.Next()
	}
}
