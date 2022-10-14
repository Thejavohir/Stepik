// package main

// import "fmt"

// func main() {

// 	var a int
// 	var b int
// 	var c int
// 	fmt.Scan(&a) // считаем переменную 'a' с консои
// 	fmt.Scan(&b) // считаем переменную 'b' с консоли

// 	a = a * a
// 	b = b * b
// 	c = a + b
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	fmt.Scan(&a)
// 	a = a * a
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	var N int
// 	fmt.Scan(&N)
// 	N = N % 10
// 	fmt.Println(N)
// }


package main

import "fmt"

const (
	DegInHour = 30
	MinInDeg  = 2
	TimeFormat = "It is %d hours %d minutes."
)

func main() {
	var input int32
	fmt.Scan(&input)

	hours := input / DegInHour
	minutes := (input % DegInHour) * MinInDeg

	fmt.Printf(TimeFormat, hours, minutes)
}