package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Struktur data untuk pesan rahasia antar negara
type PesanRahasia struct {
	ID        string    `json:"id"`
	Pengirim  string    `json:"pengirim"`
	Isi       string    `json:"isi"`
	Waktu     time.Time `json:"waktu"`
	Level     int       `json:"level_bahaya"`
}

func main() {
	fmt.Println("===========================================")
	fmt.Println("   DARAKANULA GLOBAL COMMUNICATION SYSTEM  ")
	fmt.Println("===========================================")
	
	rand.Seed(time.Now().UnixNano())
	
	// Daftar negara yang terlibat dalam percakapan kiamat
	daftarNegara := []string{"Sektor_A", "Sektor_B", "Sektor_C", "Bunker_Pusat"}

	// Simulasi ratusan baris logika komunikasi
	for i := 1; i <= 50; i++ {
		pengirim := daftarNegara[rand.Intn(len(daftarNegara))]
		level := rand.Intn(10)
		
		fmt.Printf("[LOG #%03d] Memindai sinyal dari %s...\n", i, pengirim)
		
		if level > 7 {
			fmt.Printf("   >> PERINGATAN: Sinyal Terdistorsi oleh Anomali Saraf!\n")
		} else {
			fmt.Printf("   >> Status: Komunikasi stabil.\n")
		}
		
		// Jeda sedikit agar simulasi terlihat nyata
		if i%10 == 0 {
			fmt.Println("--- Melakukan sinkronisasi data dengan Jansen (JSON) ---")
		}
	}

	// Membuat objek pesan rahasia
	pesan := PesanRahasia{
		ID:       "DRK-99",
		Pengirim: "Pusat Kimia",
		Isi:      "Gas beracun terdeteksi. Aktifkan sistem keamanan Rust segera!",
		Waktu:    time.Now(),
		Level:    9,
	}

	// Simulasi enkripsi data ke format JSON
	jsonData, _ := json.MarshalIndent(pesan, "", "  ")
	
	fmt.Println("\n[HASIL ENKRIPSI PESAN NEGARA]")
	fmt.Println(string(jsonData))

	fmt.Println("\n===========================================")
	fmt.Println("SISTEM GO BERHASIL MENYAMBUNGKAN DUNIA")
	fmt.Println("===========================================")
}
