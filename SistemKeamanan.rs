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
// SISTEM KEAMANAN BUNKER DARAKANULA (RUST CORE)
fn main() {
    let nama_sistem = "Aegis-Rust";
    let tingkat_bahaya = 95.0; // Persentase ancaman kimia
    let kode_akses = "DARAKANULA-2026";
    
    println!("--- MENGAKTIFKAN PROTOKOL {} ---", nama_sistem);
    println!("Menganalisis integritas pintu bunker...");

    let status_saraf_pusat = true;
    let kebocoran_gas_kimia = false;

    if status_saraf_pusat && !kebocoran_gas_kimia {
        println!("Status: Aman. Kode Akses {} Diterima.", kode_akses);
        println!("Pesan: Semua pintu dikunci secara otomatis dari anomali luar.");
    } else {
        println!("PERINGATAN: Kerusakan sistem saraf terdeteksi!");
        println!("Protokol Kiamat diaktifkan segera.");
    }

    // Simulasi hitungan keamanan fisik
    for i in 1..=5 {
        println!("Memindai area sektor {}... Selesai.", i);
    }

    println!("Sistem Rust berjalan stabil di balik layar.");
}
