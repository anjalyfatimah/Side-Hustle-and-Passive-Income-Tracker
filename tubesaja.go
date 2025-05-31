package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxKeterampilan = 20

type KategoriSideHustle string

const (
	PekerjaanSampingan KategoriSideHustle = "PekerjaanSampingan"
	Freelancing        KategoriSideHustle = "Freelancing"
	BisnisDigital      KategoriSideHustle = "BisnisDigital"
	InvestasiPasif     KategoriSideHustle = "InvestasiPasif"
)

type SideHustle struct {
	ID             int
	Nama           string
	Kategori       KategoriSideHustle
	MinPenghasilan float64
	MaxPenghasilan float64
	BiayaAwal      float64
	WaktuPerMinggu int
	Keterampilan   [MaxKeterampilan]string
	Deskripsi      string
}

type SumberPendapatan struct {
	Nama        string
	Penghasilan float64
}

type BiayaOperasional struct {
	Nama  string
	Biaya float64
}

func DisplayWelcome() {
	fmt.Println("====================================")
	fmt.Println("     SELAMAT DATANG DI APLIKASI     ")
	fmt.Println("        Pencari Side Hustle         ")
	fmt.Println("====================================")
	fmt.Println("Aplikasi ini membantu kamu mencari dan memilih side hustle terbaik sesuai kebutuhan.")
	fmt.Println("Silakan pilih menu untuk melanjutkan...")
	fmt.Println()
	fmt.Print("Tekan Enter untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func GetSideHustleDatabase() ([20]SideHustle, int) {
	var database [20]SideHustle
	database[0] = SideHustle{
		ID:             1,
		Nama:           "Mengajar Les Privat",
		Kategori:       PekerjaanSampingan,
		MinPenghasilan: 1000000,
		MaxPenghasilan: 4000000,
		BiayaAwal:      500000,
		WaktuPerMinggu: 8,
		Keterampilan:   [MaxKeterampilan]string{"Mengajar", "Komunikasi", "Kesabaran"},
		Deskripsi:      "Memberikan les privat untuk siswa SD, SMP, atau SMA dalam mata pelajaran tertentu",
	}
	database[1] = SideHustle{
		ID:             2,
		Nama:           "Freelance Writer",
		Kategori:       Freelancing,
		MinPenghasilan: 3000000,
		MaxPenghasilan: 8000000,
		BiayaAwal:      1000000,
		WaktuPerMinggu: 15,
		Keterampilan:   [MaxKeterampilan]string{"Menulis", "Riset", "SEO"},
		Deskripsi:      "Menulis artikel, blog, atau konten untuk klien di berbagai platform freelance",
	}
		database[2] = SideHustle{
		ID:             3,
		Nama:           "Desain Grafis Freelance",
		Kategori:       Freelancing,
		MinPenghasilan: 2500000,
		MaxPenghasilan: 10000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 12,
		Keterampilan:   [MaxKeterampilan]string{"Photoshop", "Illustrator", "Kreativitas"},
		Deskripsi:      "Membuat desain logo, poster, brosur dan materi visual lainnya untuk klien",
	}
	database[3] = SideHustle{
    ID:             4,
    Nama:           "Toko Online",
    Kategori:       BisnisDigital,
    MinPenghasilan: 2000000,
    MaxPenghasilan: 15000000,
    BiayaAwal:      5000000,
    WaktuPerMinggu: 20,
    Keterampilan:   [MaxKeterampilan]string{"Pemasaran", "Manajemen Persediaan", "Layanan Pelanggan"},
    Deskripsi:      "Menjual produk secara online melalui marketplace atau website sendiri",
	}
	database[4] = SideHustle{
    ID:             5,
    Nama:           "Dropshipping",
    Kategori:       BisnisDigital,
    MinPenghasilan: 1500000,
    MaxPenghasilan: 2000000,
    BiayaAwal:      2000000,
    WaktuPerMinggu: 15,
    Keterampilan:   [MaxKeterampilan]string{"Pemasaran Digital", "Pemilihan Produk", "Analisis Pasar"},
    Deskripsi:      "Menjual produk tanpa menyimpan stok, dengan mengirim langsung dari supplier ke pelanggan",
	}
	database[5] = SideHustle{
		ID:             6,
		Nama:           "Investasi Reksa Dana",
		Kategori:       InvestasiPasif,
		MinPenghasilan: 500000,
		MaxPenghasilan: 5000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 5,
		Keterampilan:   [MaxKeterampilan]string{"Analisis Keuangan", "Manajemen Risiko"},
		Deskripsi:      "Investasi pada instrumen reksa dana untuk mendapatkan return jangka panjang",
	}
	database[6] = SideHustle{
		ID:             7,
		Nama:           "Saham Dividen",
		Kategori:       InvestasiPasif,
		MinPenghasilan: 1000000,
		MaxPenghasilan: 20000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 4,
		Keterampilan:   [MaxKeterampilan]string{"Analisis Fundamental", "Kesabaran", "Riset"},
		Deskripsi:      "Investasi pada saham yang membagikan dividen secara rutin",
	}
	database[7] = SideHustle{
		ID:             8,
		Nama:           "YouTube Creator",
		Kategori:       BisnisDigital,
		MinPenghasilan: 1000000,
		MaxPenghasilan: 40000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 12,
		Keterampilan:   [MaxKeterampilan]string{"Editing Video", "Public Speaking", "SEO"},
		Deskripsi:      "Membuat konten video di YouTube dan mendapatkan penghasilan dari iklan dan sponsor",
	}
		database[8] = SideHustle{
		ID:             9,
		Nama:           "Ojek Online",
		Kategori:       PekerjaanSampingan,
		MinPenghasilan: 1500000,
		MaxPenghasilan: 3500000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 12,
		Keterampilan:   [MaxKeterampilan]string{"Mengemudi", "Navigasi", "Layanan Pelanggan"},
		Deskripsi:      "Menjadi pengemudi ojek online di waktu luang",
	}	

	database[9] = SideHustle{
		ID:             10,
		Nama:           "Food Delivery",
		Kategori:       PekerjaanSampingan,
		MinPenghasilan: 1000000,
		MaxPenghasilan: 15000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 20,
		Keterampilan:   [MaxKeterampilan]string{"Mengemudi", "Manajemen Waktu", "Ketelitian"},
		Deskripsi:      "Mengantar makanan melalui aplikasi delivery",
	}
	database[10] = SideHustle{
		ID:             11,
		Nama:           "Freelance Web Developer",
		Kategori:       Freelancing,
		MinPenghasilan: 5000000,
		MaxPenghasilan: 18000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 10,
		Keterampilan:   [MaxKeterampilan]string{"HTML/CSS", "JavaScript", "PHP/Python"},
		Deskripsi:      "Membuat dan mengembangkan website untuk klien",
	}
	database[11] = SideHustle{
		ID:             12,
		Nama:           "Freelance Mobile App Developer",
		Kategori:       Freelancing,
		MinPenghasilan: 6000000,
		MaxPenghasilan: 10000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 12,
		Keterampilan:   [MaxKeterampilan]string{"Java/Kotlin", "Swift", "React Native"},
		Deskripsi:      "Membuat aplikasi mobile untuk Android atau iOS",
	}
	database[12] = SideHustle{
		ID:             13,
		Nama:           "Sewa Properti",
		Kategori:       InvestasiPasif,
		MinPenghasilan: 3000000,
		MaxPenghasilan: 8000000,
		BiayaAwal:      5000000,
		WaktuPerMinggu: 20,
		Keterampilan:   [MaxKeterampilan]string{"Manajemen Properti", "Negosiasi", "Perawatan"},
		Deskripsi:      "Menyewakan rumah, apartemen, atau ruang komersial",
	}
	database[13] = SideHustle{
		ID:             14,
		Nama:           "Affiliate Marketing",
		Kategori:       BisnisDigital,
		MinPenghasilan: 1000000,
		MaxPenghasilan: 10000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 12,
		Keterampilan:   [MaxKeterampilan]string{"Pemasaran Digital", "SEO", "Copywriting"},
		Deskripsi:      "Mempromosikan produk orang lain dan mendapatkan komisi dari setiap penjualan",
	}
	database[14] = SideHustle{
		ID:             15,
		Nama:           "Fotografi",
		Kategori:       Freelancing,
		MinPenghasilan: 2000000,
		MaxPenghasilan: 7000000,
		BiayaAwal:      5000000,
		WaktuPerMinggu: 10,
		Keterampilan:   [MaxKeterampilan]string{"Fotografi", "Editing Foto", "Komunikasi"},
		Deskripsi:      "Menjadi fotografer untuk acara, potret, produk, dll",
	}
	database[15] = SideHustle{
		ID:             16,
		Nama:           "Print on Demand",
		Kategori:       BisnisDigital,
		MinPenghasilan: 1500000,
		MaxPenghasilan: 5000000,
		BiayaAwal:      3000000,
		WaktuPerMinggu: 10,
		Keterampilan:   [MaxKeterampilan]string{"Desain", "Pemasaran", "Branding"},
		Deskripsi:      "Menjual produk dengan desain kustom yang dicetak sesuai pesanan",
	}
	database[16] = SideHustle{
		ID:             17,
		Nama:           "Virtual Assistant",
		Kategori:       Freelancing,
		MinPenghasilan: 2500000,
		MaxPenghasilan: 5000000,
		BiayaAwal:      5000000,
		WaktuPerMinggu: 20,
		Keterampilan:   [MaxKeterampilan]string{"Administrasi", "Manajemen Waktu", "Teknologi"},
		Deskripsi:      "Memberikan bantuan administratif secara jarak jauh kepada klien",
	}
	database[17] = SideHustle{
		ID:             18,
		Nama:           "Transcriber",
		Kategori:       Freelancing,
		MinPenghasilan: 1500000,
		MaxPenghasilan: 15000000,
		BiayaAwal:      500000,
		WaktuPerMinggu: 15,
		Keterampilan:   [MaxKeterampilan]string{"Mendengarkan", "Mengetik Cepat", "Bahasa"},
		Deskripsi:      "Mentranskripsi rekaman audio menjadi teks tertulis",
	}
	database[18] = SideHustle{
		ID:             19,
		Nama:           "Course Creator",
		Kategori:       BisnisDigital,
		MinPenghasilan: 2000000,
		MaxPenghasilan: 10000000,
		BiayaAwal:      2000000,
		WaktuPerMinggu: 20,
		Keterampilan:   [MaxKeterampilan]string{"Mengajar", "Produksi Video", "Pemasaran"},
		Deskripsi:      "Membuat dan menjual kursus online di platform seperti Udemy atau Skillshare",
	}
	database[19] = SideHustle{
		ID:             20,
		Nama:           "P2P Lending",
		Kategori:       InvestasiPasif,
		MinPenghasilan: 500000,
		MaxPenghasilan: 30000000,
		BiayaAwal:      1000000,
		WaktuPerMinggu: 5,
		Keterampilan:   [MaxKeterampilan]string{"Analisis Kredit", "Manajemen Risiko"},
		Deskripsi:      "Meminjamkan uang melalui platform peer-to-peer lending",
	}

	return database, 20
}

func TampilkanDetailSideHustle(sh SideHustle) {
	fmt.Println("\n=== DETAIL SIDE HUSTLE ===")
	fmt.Printf("ID               : %d\n", sh.ID)
	fmt.Printf("Nama             : %s\n", sh.Nama)
	fmt.Printf("Kategori         : %s\n", sh.Kategori)
	fmt.Printf("Estimasi Penghasilan: Rp %.0f - Rp %.0f per bulan\n", sh.MinPenghasilan, sh.MaxPenghasilan)
	fmt.Printf("Biaya Awal       : Rp %.0f\n", sh.BiayaAwal)
	fmt.Printf("Waktu yang Dibutuhkan: %d jam per minggu\n", sh.WaktuPerMinggu)

	var keterampilanArr [MaxKeterampilan]string
	jumlah := 0
	for i := 0; i < MaxKeterampilan; i++ {
		if sh.Keterampilan[i] != "" {
			keterampilanArr[jumlah] = sh.Keterampilan[i]
			jumlah++
		}
	}

	keterampilanGabungan := ""
	for i := 0; i < jumlah; i++ {
		keterampilanGabungan += keterampilanArr[i]
		if i < jumlah-1 {
			keterampilanGabungan += ", "
		}
	}

	fmt.Printf("Keterampilan yang Dibutuhkan: %s\n", keterampilanGabungan)
	fmt.Printf("Deskripsi        : %s\n", sh.Deskripsi)
}

func CariSideHustle(database [20]SideHustle, total int, namaKunci string, kategori KategoriSideHustle, minPenghasilan, maxPenghasilan float64, maxWaktu int, maxBiaya float64, keterampilanKunci string) ([100]SideHustle, int) {
	var hasil [100]SideHustle
	jumlah := 0
	for i := 0; i < total; i++ {
		sh := database[i]
		valid := true

		if namaKunci != "" && !strings.Contains(strings.ToLower(sh.Nama), strings.ToLower(namaKunci)) {
			valid = false
		}
		if valid && kategori != "" && sh.Kategori != kategori {
			valid = false
		}
		if valid && minPenghasilan > 0 && sh.MaxPenghasilan < minPenghasilan {
			valid = false
		}
		if valid && maxPenghasilan > 0 && sh.MinPenghasilan > maxPenghasilan {
			valid = false
		}
		if valid && maxWaktu > 0 && sh.WaktuPerMinggu > maxWaktu {
			valid = false
		}
		if valid && maxBiaya > 0 && sh.BiayaAwal > maxBiaya {
			valid = false
		}
		if valid && keterampilanKunci != "" {
			keterampilanDitemukan := false
			for k := 0; k < MaxKeterampilan; k++ {
				kat := sh.Keterampilan[k]
				if kat != "" && strings.Contains(strings.ToLower(kat), strings.ToLower(keterampilanKunci)) {
					keterampilanDitemukan = true
				}
			}
			if !keterampilanDitemukan {
				valid = false
			}
		}

		if valid && jumlah < 100 {
			hasil[jumlah] = sh
			jumlah++
		}
	}
	return hasil, jumlah
}

func TampilkanHasilPencarian(hasil [100]SideHustle, n int) {
	if n == 0 {
		fmt.Println("Tidak ditemukan side hustle yang sesuai dengan kriteria pencarian")
		return
	}

	fmt.Println("\n=== HASIL PENCARIAN ===")
	fmt.Printf("%-3s  %-30s  %-20s  %-35s  %-15s  %-12s\n", "ID", "Nama", "Kategori", "Penghasilan/Bulan", "Modal Awal", "Waktu/Minggu")
	fmt.Println(strings.Repeat("=", 115))

	for i := 0; i < n; i++ {
		sh := hasil[i]
		fmt.Printf("%-3d  %-30s  %-20s  Rp %10.0f - Rp %-10.0f  Rp %13.0f  %10d jam\n",
			sh.ID, sh.Nama, sh.Kategori,
			sh.MinPenghasilan, sh.MaxPenghasilan, sh.BiayaAwal, sh.WaktuPerMinggu)
	}
}

func SortSideHustleByMinPenghasilan(arr *[20]SideHustle, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j].MinPenghasilan > arr[j+1].MinPenghasilan {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func CariPenghasilanTertinggi(database [20]SideHustle, total int) SideHustle {
	maxSH := database[0]
	for i := 1; i < total; i++ {
		if database[i].MaxPenghasilan > maxSH.MaxPenghasilan {
			maxSH = database[i]
		}
	}
	return maxSH
}

func CariPenghasilanTerendah(database [20]SideHustle, total int) SideHustle {
	minSH := database[0]
	for i := 1; i < total; i++ {
		if database[i].MinPenghasilan < minSH.MinPenghasilan {
			minSH = database[i]
		}
	}
	return minSH
}

func InputSumberPendapatan() ([100]SumberPendapatan, int) {
	var sumberPendapatan [100]SumberPendapatan
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	stopInput := false

	for !stopInput {
		if count >= 100 {
			fmt.Println("Jumlah input sumber pendapatan maksimal 100.")
			stopInput = true
		} else {
			fmt.Print("Nama sumber pendapatan (kosongkan untuk selesai): ")
			scanner.Scan()
			nama := scanner.Text()
			if nama == "" {
				stopInput = true
			} else {
				fmt.Print("Jumlah penghasilan (Rp): ")
				scanner.Scan()
				penghasilanStr := scanner.Text()
				penghasilan, err := strconv.ParseFloat(penghasilanStr, 64)
				if err != nil || penghasilan < 0 {
					fmt.Println("Input penghasilan harus angka positif, silakan coba lagi")
				} else {
					sumberPendapatan[count] = SumberPendapatan{Nama: nama, Penghasilan: penghasilan}
					count++
				}
			}
		}
	}
	return sumberPendapatan, count
}

func InputBiayaOperasional() ([100]BiayaOperasional, int) {
	var biayaOperasional [100]BiayaOperasional
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	stopInput := false

	for !stopInput {
		if count >= 100 {
			fmt.Println("Jumlah input biaya operasional maksimal 100.")
			stopInput = true
		} else {
			fmt.Print("Nama biaya operasional (kosongkan untuk selesai): ")
			scanner.Scan()
			nama := scanner.Text()
			if nama == "" {
				stopInput = true
			} else {
				fmt.Print("Jumlah biaya (Rp): ")
				scanner.Scan()
				biayaStr := scanner.Text()
				biaya, err := strconv.ParseFloat(biayaStr, 64)
				if err != nil || biaya < 0 {
					fmt.Println("Input biaya harus angka positif, silakan coba lagi")
				} else {
					biayaOperasional[count] = BiayaOperasional{Nama: nama, Biaya: biaya}
					count++
				}
			}
		}
	}
	return biayaOperasional, count
}

func HitungPenghasilanDanKeuntungan(sumberPendapatan [100]SumberPendapatan, sumberCount int, biayaOperasional [100]BiayaOperasional, biayaCount int) {
	var totalPendapatan float64
	var totalBiaya float64

	fmt.Println("\n=== RINGKASAN KEUANGAN ===")

	fmt.Println("\nSUMBER PENDAPATAN:")
	fmt.Printf("%-30s %-15s\n", "Nama", "Jumlah (Rp)")
	fmt.Println(strings.Repeat("-", 45))
	for i := 0; i < sumberCount; i++ {
		p := sumberPendapatan[i]
		fmt.Printf("%-30s Rp %-15.0f\n", p.Nama, p.Penghasilan)
		totalPendapatan += p.Penghasilan
	}
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%-30s Rp %-15.0f\n", "TOTAL PENDAPATAN", totalPendapatan)

	fmt.Println("\nBIAYA OPERASIONAL:")
	fmt.Printf("%-30s %-15s\n", "Nama", "Jumlah (Rp)")
	fmt.Println(strings.Repeat("-", 45))
	for i := 0; i < biayaCount; i++ {
		b := biayaOperasional[i]
		fmt.Printf("%-30s Rp %-15.0f\n", b.Nama, b.Biaya)
		totalBiaya += b.Biaya
	}
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%-30s Rp %-15.0f\n", "TOTAL BIAYA", totalBiaya)

	keuntungan := totalPendapatan - totalBiaya
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%-30s Rp %-15.0f\n", "KEUNTUNGAN/KERUGIAN", keuntungan)

	if keuntungan > 0 {
		fmt.Println("\nStatus: UNTUNG")
		fmt.Printf("Margin keuntungan: %.2f%%\n", (keuntungan/totalPendapatan)*100)
	} else if keuntungan < 0 {
		fmt.Println("\nStatus: RUGI")
		fmt.Printf("Margin kerugian: %.2f%%\n", (keuntungan/totalPendapatan)*-100)
	} else {
		fmt.Println("\nStatus: IMPAS (Break Even)")
	}
}

func TampilkanSemuaSideHustle(database [20]SideHustle, total int) {
	fmt.Println("\n=== SEMUA SIDE HUSTLE ===")
	fmt.Printf("%-3s %-30s %-20s %-35s %-15s %-12s\n",
		"ID", "Nama", "Kategori", "Penghasilan/Bulan", "Modal Awal", "Waktu/Minggu")
	fmt.Println(strings.Repeat("=", 125))

	for i := 0; i < total; i++ {
		sh := database[i]
		fmt.Printf("%-3d %-30s %-20s Rp %10.0f - Rp %-10.0f Rp %13.0f %11d jam\n",
			sh.ID, sh.Nama, sh.Kategori, sh.MinPenghasilan, sh.MaxPenghasilan, sh.BiayaAwal, sh.WaktuPerMinggu)
	}
}

func main() {
	DisplayWelcome()
	database, totalData := GetSideHustleDatabase()
	scanner := bufio.NewScanner(os.Stdin)
	exitProgram := false

	for !exitProgram {
		fmt.Println("\n===== PENCARI SIDE HUSTLE =====")
		fmt.Println("1. Tampilkan Semua Side Hustle")
		fmt.Println("2. Cari Side Hustle (Filter Lanjutan)")
		fmt.Println("3. Lihat Detail Side Hustle (Cari by ID)")
		fmt.Println("4. Kalkulator Penghasilan dan Keuntungan")
		fmt.Println("5. Tampilkan Side Hustle Terurut (Min Penghasilan Ascending)")
		fmt.Println("6. Cari Side Hustle dengan Penghasilan Tertinggi atau Terendah")
		fmt.Println("0. Keluar")
		fmt.Print("\nPilih menu: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "0":
			fmt.Println("Terima kasih telah menggunakan Pencari Side Hustle!")
			exitProgram = true

		case "1":
				TampilkanSemuaSideHustle(database, totalData)

		case "2":
			fmt.Print("Masukkan kata kunci nama (kosongkan jika tidak pakai): ")
			scanner.Scan()
			namaKunci := scanner.Text()

			fmt.Print("Masukkan kategori (PekerjaanSampingan/Freelancing/BisnisDigital/InvestasiPasif) (kosongkan jika tidak pakai): ")
			scanner.Scan()
			kategoriInput := scanner.Text()
			var kategori KategoriSideHustle
			if kategoriInput == "PekerjaanSampingan" {
				kategori = PekerjaanSampingan
			} else if kategoriInput == "Freelancing" {
				kategori = Freelancing
			} else if kategoriInput == "BisnisDigital" {
				kategori = BisnisDigital
			} else if kategoriInput == "InvestasiPasif" {
				kategori = InvestasiPasif
			} else {
				kategori = ""
			}

			fmt.Print("Masukkan minimal penghasilan (0 jika tidak pakai): ")
			scanner.Scan()
			minPenghasilanStr := scanner.Text()
			minPenghasilan, err := strconv.ParseFloat(minPenghasilanStr, 64)
			if err != nil {
				fmt.Println("Input tidak valid, masukkan angka.")
				minPenghasilan = 0
			}

			fmt.Print("Masukkan maksimal penghasilan (0 jika tidak pakai): ")
			scanner.Scan()
			maxPenghasilanStr := scanner.Text()
			maxPenghasilan, err := strconv.ParseFloat(maxPenghasilanStr, 64)
			if err != nil {
				fmt.Println("Input tidak valid, masukkan angka.")
				maxPenghasilan = 0
			}

			fmt.Print("Masukkan maksimal waktu per minggu (jam) (0 jika tidak pakai): ")
			scanner.Scan()
			maxWaktuStr := scanner.Text()
			maxWaktu, err := strconv.Atoi(maxWaktuStr)
			if err != nil {
				fmt.Println("Input tidak valid, masukkan angka.")
				maxWaktu = 0
			}

			fmt.Print("Masukkan maksimal biaya awal (0 jika tidak pakai): ")
			scanner.Scan()
			maxBiayaStr := scanner.Text()
			maxBiaya, err := strconv.ParseFloat(maxBiayaStr, 64)
			if err != nil {
				fmt.Println("Input tidak valid, masukkan angka.")
				maxBiaya = 0
			}
			
			fmt.Print("Masukkan kata kunci keterampilan (kosongkan jika tidak pakai): ")
			scanner.Scan()
			keterampilanKunci := scanner.Text()

			hasil, jumlah := CariSideHustle(database, totalData, namaKunci, kategori, minPenghasilan, maxPenghasilan, maxWaktu, maxBiaya, keterampilanKunci)

			TampilkanHasilPencarian(hasil, jumlah)

		case "3":
			fmt.Print("Masukkan ID side hustle yang ingin dilihat detailnya: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Input tidak valid")
			} else {
				found := false
				for i := 0; i < totalData; i++ {
					if database[i].ID == id {
						TampilkanDetailSideHustle(database[i])
						found = true
					}
				}
				if !found {
					fmt.Println("Side hustle dengan ID tersebut tidak ditemukan")
				}
			}

		case "4":
			sumberPendapatan, sumberCount := InputSumberPendapatan()
			biayaOperasional, biayaCount := InputBiayaOperasional()
			HitungPenghasilanDanKeuntungan(sumberPendapatan, sumberCount, biayaOperasional, biayaCount)

		case "5":
			SortSideHustleByMinPenghasilan(&database, totalData)
			fmt.Println("\nSide Hustle Terurut berdasarkan Min Penghasilan (Ascending):")
			for i := 0; i < totalData; i++ {
				TampilkanDetailSideHustle(database[i])
			}

		case "6":
			fmt.Println("\nPilih opsi pencarian:")
			fmt.Println("1. Penghasilan Tertinggi")
			fmt.Println("2. Penghasilan Terendah")
			fmt.Print("Masukkan pilihan (1-2): ")
			scanner.Scan()
			subPilihan := scanner.Text()

			if subPilihan == "1" {
				sh := CariPenghasilanTertinggi(database, totalData)
				fmt.Println("\nSide Hustle dengan Penghasilan Tertinggi:")
				TampilkanDetailSideHustle(sh)
			} else if subPilihan == "2" {
				sh := CariPenghasilanTerendah(database, totalData)
				fmt.Println("\nSide Hustle dengan Penghasilan Terendah:")
				TampilkanDetailSideHustle(sh)
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
