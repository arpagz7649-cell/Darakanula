// SISTEM KONTROLER & KAMERA 3D DARAKANULA
const gameContainer = document.body;

// Logika nengok kanan-kiri (Mouse Look)
let mouseX = 0, mouseY = 0;
document.addEventListener('mousemove', (e) => {
    if (document.pointerLockElement === gameContainer) {
        mouseX += e.movementX;
        mouseY += e.movementY;
        console.log(`Kamera Nengok ke X: ${mouseX}, Y: ${mouseY}`);
        // Di sini nanti kita hubungkan ke rotasi kamera 3D
    }
});

// Klik layar buat aktifin mode kamera game
gameContainer.addEventListener('click', () => {
    gameContainer.requestPointerLock();
    console.log("Kamera Terkunci - Mode Controller Aktif");
});

// Logika Jalan (Controller)
window.addEventListener('keydown', (e) => {
    switch(e.key.toLowerCase()) {
        case 'w': console.log("Maju ke depan..."); break;
        case 's': console.log("Mundur..."); break;
        case 'a': console.log("Geser Kiri..."); break;
        case 'd': console.log("Geser Kanan..."); break;
    }
});

console.log("Sistem Saraf JavaScript Siap Menjalankan C++, Go, dan Rust.");
