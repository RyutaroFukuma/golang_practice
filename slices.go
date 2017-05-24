package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4] //配列の1番目から(4-1)番目を取り出す
    fmt.Println(s)
}

