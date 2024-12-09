package main

import "fmt"
func main() {
    var angkaAwal, angkaAkhir, kelipatan int
    fmt.Scan(&angkaAwal, &angkaAkhir, &kelipatan)
    for i := angkaAwal; i <= angkaAkhir; i++ {
        if i % kelipatan == 0 {
            fmt.Print(i, " ")
        }
    }
}
