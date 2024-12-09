package main

import "fmt"
func main() {
    var angka int; fmt.Scan(&angka)
    // jika angka tidak habis dibagi dua
    if angka % 2 == 1 {
        // maka dia ganjil
        fmt.Print("ganjil")
    // tapi jika angka habis dibagi dua
    } else if angka % 2 == 0 {
        // maka dia genap
        fmt.Print("genap")
    }
}
