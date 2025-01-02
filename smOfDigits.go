package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&n)

    s := 0
    for n != 0 {
        s += n % 10
        n /= 10
    }

    fmt.Println("The sum of the digits is:", s)
}
