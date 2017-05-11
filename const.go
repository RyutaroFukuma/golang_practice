//定数

package main

import "fmt"

func main() {
    const name = "fkm"
    fmt.Println(name)

    const (
        sun = iota // 0
        mon // 1
        tue // 2
    )

    fmt.Println(sun, mon, tue)
}
