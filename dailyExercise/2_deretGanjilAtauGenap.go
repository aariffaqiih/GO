// package main

// import "fmt"
// func main () {
// 	var angka1, angka2 int
// 	fmt.Scan(&angka1, &angka2)
//     fmt.Print("ini adalah deret ganjil: ")
// 	for i := angka1; i <= angka2; i++ {
//         if i % 2 == 1 {
//             fmt.Print(i, " ")
//         }
// 	}
//     fmt.Println() // ini cuma buat enter
//     fmt.Print("ini adalah deret genap : ")
//     for i := angka1; i <= angka2; i++ {
//         if i % 2 == 0 {
//             fmt.Print(i, " ")
//         }
// 	}
// }

package main

import "fmt"

func main() {
    var angka1, angka2 int
    var pilihan string

    // Input angka awal dan akhir
    fmt.Print("Masukkan angka awal: ")
    fmt.Scan(&angka1)
    fmt.Print("Masukkan angka akhir: ")
    fmt.Scan(&angka2)
    fmt.Print("Pilih 'ganjil' atau 'genap': ")
    fmt.Scan(&pilihan)

    // Tentukan berdasarkan pilihan
    switch pilihan {
    case "ganjil":
        fmt.Print("Ini adalah deret ganjil: ")
        for i := angka1; i <= angka2; i++ {
            if i%2 == 1 {
                fmt.Print(i, " ")
            }
        }
    case "genap":
        fmt.Print("Ini adalah deret genap: ")
        for i := angka1; i <= angka2; i++ {
            if i%2 == 0 {
                fmt.Print(i, " ")
            }
        }
    default:
        fmt.Println("Pilihan tidak valid. Harus 'ganjil' atau 'genap'.")
    }
}
