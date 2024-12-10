package main

import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    if a < b && b < c {
        fmt.Print("angka terbesar adalah: ",c)
    } else if a < c && c < b {
        fmt.Print("angka terbesar adalah: ",b)
    } else if c < a && a < b {
        fmt.Print("angka terbesar adalah: ",b)
    }else if b < c && c < a {
        fmt.Print("angka terbesar adalah: ",a)
    }else if c < b && b < a {
        fmt.Print("angka terbesar adalah: ",a)
    }else if b < a && a < c {
        fmt.Print("angka terbesar adalah: ",c)
    }
}

// package main

// import "fmt"

// func main() {
//     var a, b, c int
//     fmt.Scan(&a, &b, &c)
    
//     // Mulai dengan asumsi 'a' sebagai angka terbesar
//     terbesar := a
    
//     // Bandingkan dengan 'b'
//     if b > terbesar {
//         terbesar = b
//     }
    
//     // Bandingkan dengan 'c'
//     if c > terbesar {
//         terbesar = c
//     }
    
//     fmt.Print("Angka terbesar adalah: ", terbesar)
// }
