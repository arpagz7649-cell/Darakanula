fn main() {
    // Logika kimia dan keamanan bunker
    let tingkat_radiasi: f32 = 89.5;
    let status_oksidasi = "Kritis - Logam Berkarat";
    let bunker_terkunci = true;

    println!("=== PROTOKOL RUST (KARAT) DIAKTIFKAN ===");
    println!("Mendeteksi anomali reaksi kimia di udara...");
    println!("Tingkat Radiasi: {} %", tingkat_radiasi);
    println!("Status Lingkungan: {}", status_oksidasi);
    
    if bunker_terkunci {
        println!("Sistem Saraf Aman. Pintu utama bunker berhasil dikunci dari entitas luar.");
    } else {
        println!("PERINGATAN: Kebocoran terdeteksi!");
    }
}
