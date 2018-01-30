package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println("Hello")

	fmt.Println("Go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "shorting"
	fmt.Println(f)

	fmt.Println(s)

	const n = 500000000
	const dd = 3e20 / n

	fmt.Println(dd)
	fmt.Println(int64(dd))

	fmt.Println(math.Sin(n))
}
