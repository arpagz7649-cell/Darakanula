package main

import "fmt"

func main() {
    // Simulasi pengiriman pesan rahasia antar negara
    statusNegara := "SIAGA 1"
    pesan := "Anomali terdeteksi di sektor saraf pusat. Segera aktifkan protokol Darakanula!"

    fmt.Println("--- SISTEM KOMUNIKASI GO AKTIF ---")
    fmt.Printf("Status Global: %s\n", statusNegara)
    fmt.Println("Mengirim pesan rahasia...")
    fmt.Printf("Isi Pesan: %s\n", pesan)
    fmt.Println("Pesan berhasil dienkripsi dan dikirim.")
}

