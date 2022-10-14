// package main

// import "fmt"

// func main() {
//     var n int
//     fmt.Scanln(&n)
//     if n>0 {
//         fmt.Println("Число положительное")
// 	} else if n<0 {
//             fmt.Println("Число отрицательное")
// 	} else{
//             fmt.Println("Ноль")
// 	}
//  }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	a := n % 10
// 	b := n / 10 % 10
// 	c := n / 100
// 	if a != b && a != c && b != c {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	var a int
// 	fmt.Scan(&n)
// 	count := 0
// 	a = n
// 	for a > 0 {
// 		a = a / 10
// 		count++
// 	}
// 	if count == 1 {
// 		fmt.Println(n)
// 	} else if count == 2 {
// 		fmt.Println(n / 10)
// 	} else if count == 3 {
// 		fmt.Println(n / 100)
// 	} else if count == 4 {
// 		fmt.Println(n / 1000)
// 	} else if count == 5 {
// 		fmt.Println(n / 10000)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	var sum1 int
// 	var sum2 int
// 	var a int
// 	fmt.Scan(&n)
// 	count := 0
// 	a = n
// 	for a > 0 {
// 		a = a / 10
// 		count++
// 	}
// 	if count == 6 {
// 		count = 0
// 		for n > 0 {
// 			if count < 3 {
// 				b := n % 10
// 				sum2 += b
// 				n = n / 10
// 				count++
// 			} else {
// 				c := n % 10
// 				sum1 += c
// 				n = n / 10
// 			}
// 		}
// 	}
// 	if sum1 == sum2 {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%400 == 0 || n%4 == 0 && n%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
