package main

import (
    "fmt"
)

func ExchangeCoin(amount int) []int {
    // Daftar pecahan koin yang tersedia
    coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
    result := []int{}

    // Jika jumlahnya 0, kembalikan slice kosong
    if amount == 0 {
        return result
    }

    // Proses pembagian koin
    for _, coin := range coins {
        // Selama sisa amount masih lebih besar atau sama dengan pecahan koin
        for amount >= coin {
            result = append(result, coin) // Tambahkan koin ke hasil
            amount -= coin                 // Kurangi jumlah dengan pecahan koin
        }
    }

    return result
}

func main() {
    // Contoh test case
    fmt.Println(ExchangeCoin(1752)) // Output: [1000 500 200 50 1 1]
    fmt.Println(ExchangeCoin(5000)) // Output: [1000 1000 1000 1000 1000]
    fmt.Println(ExchangeCoin(1234)) // Output: [1000 200 20 10 1 1 1 1]
    fmt.Println(ExchangeCoin(0))    // Output: []
}
