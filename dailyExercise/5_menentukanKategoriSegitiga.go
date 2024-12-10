// package main

// import "fmt"
// func main() {
//     var a, b, c int
//     fmt.Scan(&a, &b, &c)
//     if a + b <= c || b + c <= a || c + a <= b || b + a <= c || c + b <= a || a + c <= b {
//         fmt.Print("sisi-sisi itu tidak bisa membuat segitiga")
//     } else if a == b && b != c || a == c && c != b || b == c && c != a {
//         fmt.Print("segitiga sama kaki")
//     } else if a == b && b == c {
//         fmt.Print("segitiga sama sisi")
//     } else if a != b && b != c {
//         fmt.Print("segitiga sembarang")
//     }
// }

package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    // Validasi apakah sisi-sisi dapat membentuk segitiga
    if a+b <= c || b+c <= a || c+a <= b {
        fmt.Println("Sisi-sisi itu tidak bisa membuat segitiga")
    } else if a == b && b == c {
        fmt.Println("Segitiga sama sisi")
    } else if a == b || b == c || a == c {
        fmt.Println("Segitiga sama kaki")
    } else {
        fmt.Println("Segitiga sembarang")
    }
}
