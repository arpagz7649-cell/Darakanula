// Fungsi ini dipanggil saat tombol "MASUK KE AREA" ditekan di HTML
function masukGame() {
    // 1. Ubah tulisan tombol biar ada efek "loading" horor
    const btn = document.getElementById('enter-btn');
    btn.innerText = "MENGHUBUNGKAN...";
    btn.style.color = "#444";
    btn.style.borderColor = "#444";
    btn.style.pointerEvents = "none"; // Biar gak bisa diklik 2x

    // 2. Simulasi menyadap jaringan (Skenario Kiamat)
    console.log("[SYSTEM] Membuka gerbang...");
    console.log("[INTELLIGENCE] Menyadap percakapan antar negara...");
    console.log("[WARNING] Sinyal anomali terdeteksi di koordinatmu.");

    // 3. Pindah layar setelah delay 2 detik (biar tegang)
    setTimeout(() => {
        // Sembunyikan Lobby
        document.getElementById('lobby').style.display = 'none';
        
        // Munculkan Ruangan Game
        document.getElementById('game-room').style.display = 'block';

        // Panggil fungsi untuk memulai simulasi
        inisialisasiDunia();
    }, 2000);
}

// Fungsi untuk menyiapkan dunia game setelah masuk
function inisialisasiDunia() {
    console.log("[SYSTEM] Area Game Dimuat. Fisika dan Biologi Aktif.");
    
    // Nanti di sini kita akan panggil logika Python atau mekanik Hantu-nya
    alert("Koneksi berhasil. Kamu sekarang sendirian di area ini. Jangan menoleh ke belakang.");
    
    // Ubah background jadi lebih gelap pas masuk game
    document.body.style.backgroundColor = "#020000";
}

