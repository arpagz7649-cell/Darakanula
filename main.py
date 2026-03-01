import random
import time
import json

class OtakGame:
    def __init__(self):
        # Memuat data dari JSON yang kita bahas tadi
        self.data_skenario = {
            "status_kiamat": "Fase Awal",
            "deteksi_hantu": False,
            "tingkat_bahaya": 0
        }
        self.posisi_hantu = 0 # 0 = Jauh, 100 = Tepat di belakangmu

    def hitung_fisika_hantu(self, gerakan_pemain):
        # Logika Fisika: Semakin pemain berisik, hantu semakin cepat mendekat
        if gerakan_pemain == "lari":
            self.posisi_hantu += random.randint(10, 20)
        else:
            self.posisi_hantu += random.randint(1, 5)
            
        return self.posisi_hantu

    def cek_biologi_ketakutan(self, waktu_diam):
        # Logika Biologi: Jika terlalu lama di gelap, tingkat bahaya naik
        if waktu_diam > 10:
            self.data_skenario["tingkat_bahaya"] += 1
        return self.data_skenario["tingkat_bahaya"]

    def siarkan_percakapan_negara(self):
        percakapan = [
            "NEGARA 1: Objek terlihat di radar.",
            "NEGARA 2: Jangan biarkan dia menyeberang perbatasan!",
            "SISTEM: Sinyal terputus... Anomali terdeteksi."
        ]
        return random.choice(percakapan)

# Menjalankan simulasi sederhana
game = OtakGame()
print("--- SISTEM HANTU AKTIF ---")
print(game.siarkan_percakapan_negara())
