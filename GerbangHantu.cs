using System;

namespace GameHantu
{
    class GerbangHantu
    {
        // Variabel untuk status saraf dan lingkungan
        public bool isGameStarted = false;
        public string statusDunia = "Normal";

        // Fungsi yang dipanggil saat tombol MASUK ditekan
        public void MasukKeArea()
        {
            isGameStarted = true;
            statusDunia = "Anomali Terdeteksi";
            
            Console.WriteLine("--- SISTEM C# (C PAGAR) AKTIF ---");
            Console.WriteLine("Status: " + statusDunia);
            Console.WriteLine("Log: Menghubungkan saraf pemain ke simulasi...");
            
            MulaiSkenarioKiamat();
        }

        void MulaiSkenarioKiamat()
        {
            // Logika Skenario: Menyiapkan percakapan antar negara
            string[] percakapan = {
                "Sinyal Radio: Unit 01, segera evakuasi!",
                "Sinyal Radio: Objek tidak terlihat di mata telanjang...",
                "Sinyal Radio: Saraf sensorik mulai terganggu."
            };
            
            Console.WriteLine(percakapan[1]);
        }
    }
}
