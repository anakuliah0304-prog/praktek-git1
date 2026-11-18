
package main

import "fmt"

func main() {
    const TARIF_PER_KG = 5500
    const BIAYA_ANTAR = 3000

    var berat float64
    var uangDiterima int

    fmt.Println("===== PROGRAM LAUNDRY =====")

    fmt.Print("Masukkan berat cucian (kg): ")
    fmt.Scan(&berat)

    fmt.Print("Masukkan uang diterima: Rp")
    fmt.Scan(&uangDiterima)

    biayaCuci := int(berat*float64(TARIF_PER_KG) + 0.5)

    total := biayaCuci + BIAYA_ANTAR

    totalBayar := ((total + 99) / 100) * 100

    kembalian := uangDiterima - totalBayar

    fmt.Println("\n===== STRUK PARKIR =====")
    fmt.Printf("Berat cucian  : %.1f kg\n", berat)
    fmt.Printf("Tarif per kg  : Rp%d\n", TARIF_PER_KG)
    fmt.Printf("Biaya cuci    : Rp%d\n", biayaCuci)
    fmt.Printf("Biaya antar   : Rp%d\n", BIAYA_ANTAR)
    fmt.Printf("Total         : Rp%d\n", total)
    fmt.Printf("Total bayar   : Rp%d\n", totalBayar)
    fmt.Printf("Uang diterima : Rp%d\n", uangDiterima)
    fmt.Printf("Kembalian     : Rp%d\n", kembalian)
}
