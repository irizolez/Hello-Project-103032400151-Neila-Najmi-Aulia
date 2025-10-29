package main

import (
	"fmt"
)

const NMAX int = 99

type workout struct {
	nama, gender, tujuan, latihan, aktivitas, hari string
	umur, frekuensi, rep, set, kalori, protein     int
	tinggi, berat, beban, jamtidur                 float64
}

type tabWorkout [NMAX]workout

func main() {
	var data tabWorkout
	var nData int
	var pilihan int

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Input Data Profil")
		fmt.Println("2. Tampilkan Data Profil")
		fmt.Println("3. Tambah Latihan")
		fmt.Println("4. Tampilkan Semua Latihan")
		fmt.Println("5. Lihat History Latihan per Hari")
		fmt.Println("6. Tambah Tracking")
		fmt.Println("7. Tampilkan Tracking")
		fmt.Println("8. Tampilkan Data Lengkap per Hari")
		fmt.Println("9. Urutkan Hari by Beban (Descending)")
		fmt.Println("10. Urutkan Hari by Beban (Ascending)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan jumlah data: ")
			fmt.Scan(&nData)
			bacaProfil(&data, nData)
		case 2:
			tampilkanProfil(data, nData)
		case 3:
			tambahLatihan(&data, nData)
		case 4:
			tampilSemuaLatihan(data, nData)
		case 5:
			lihatHistoryLatihan(data, nData)
		case 6:
			tambahTracking(&data, nData)
		case 7:
			tampilkanTracking(data, nData)
		case 8:
			pilihHari(data, nData)
		case 9:
			urutkanHariByBebanDescending(data, nData)
		case 10:
			urutkanHariByBebanAscending(data, nData)
		case 0:
			fmt.Println("Terima kasih telah menggunakan program ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func bacaProfil(A *tabWorkout, n int) {
	fmt.Println("\n=== Input Data Profil ===")
	for i := 0; i < n; i++ {
		fmt.Printf("\nData ke-%d\n", i+1)
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&A[i].nama)

		fmt.Print("Masukkan gender (P/L): ")
		fmt.Scan(&A[i].gender)

		fmt.Print("Masukkan tujuan (Power/Improve): ")
		fmt.Scan(&A[i].tujuan)

		fmt.Print("Masukkan aktivitas (Santai/Sibuk): ")
		fmt.Scan(&A[i].aktivitas)

		fmt.Print("Masukkan tinggi badan (cm): ")
		fmt.Scan(&A[i].tinggi)

		fmt.Print("Masukkan berat badan (kg): ")
		fmt.Scan(&A[i].berat)

		fmt.Print("Masukkan umur: ")
		fmt.Scan(&A[i].umur)

		fmt.Println("Profil berhasil disimpan!")
	}
}

func tampilkanProfil(A tabWorkout, n int) {
	fmt.Println("\n=== Data Profil Pengguna ===")
	for i := 0; i < n; i++ {
		fmt.Printf("\nProfil ke-%d\n", i+1)
		fmt.Printf("Nama: %s\n", A[i].nama)
		fmt.Printf("Gender: %s\n", A[i].gender)
		fmt.Printf("Tujuan: %s\n", A[i].tujuan)
		fmt.Printf("Aktivitas: %s\n", A[i].aktivitas)
		fmt.Printf("Tinggi: %.1f cm\n", A[i].tinggi)
		fmt.Printf("Berat: %.1f kg\n", A[i].berat)
		fmt.Printf("Umur: %d tahun\n", A[i].umur)
	}
}

func tambahLatihan(A *tabWorkout, n int) {
	fmt.Println("\n=== Tambah Latihan ===")
	for i := 0; i < n; i++ {
		fmt.Printf("\nUntuk %s\n", A[i].nama)

		// Pilih hari
		fmt.Println("Pilih hari latihan:")
		A[i].hari = pilihHariInput()

		fmt.Print("Masukkan jenis latihan: ")
		fmt.Scan(&A[i].latihan)

		fmt.Print("Masukkan jumlah repetisi: ")
		fmt.Scan(&A[i].rep)

		fmt.Print("Masukkan jumlah set: ")
		fmt.Scan(&A[i].set)

		fmt.Print("Masukkan beban (kg): ")
		fmt.Scan(&A[i].beban)

		fmt.Printf("Latihan berhasil ditambahkan untuk hari %s!\n", A[i].hari)
	}
}

func tampilSemuaLatihan(A tabWorkout, n int) {
	fmt.Println("\n=== Semua Data Latihan ===")
	for i := 0; i < n; i++ {
		if A[i].latihan != "" {
			fmt.Printf("\n%s - Hari %s\n", A[i].nama, A[i].hari)
			fmt.Printf("Latihan: %s\n", A[i].latihan)
			fmt.Printf("Repetisi: %d\n", A[i].rep)
			fmt.Printf("Set: %d\n", A[i].set)
			fmt.Printf("Beban: %.1f kg\n", A[i].beban)
		}
	}
}

func lihatHistoryLatihan(A tabWorkout, n int) {
	fmt.Println("\n=== Lihat History Latihan ===")
	hari := pilihHariInput()

	fmt.Printf("\nHistory Latihan Hari %s:\n", hari)
	found := false

	for i := 0; i < n; i++ {
		if A[i].hari == hari && A[i].latihan != "" {
			found = true
			fmt.Printf("\n%s\n", A[i].nama)
			fmt.Printf("Latihan: %s\n", A[i].latihan)
			fmt.Printf("Rep: %d x Set: %d\n", A[i].rep, A[i].set)
			fmt.Printf("Beban: %.1f kg\n", A[i].beban)
		}
	}

	if !found {
		fmt.Printf("Tidak ada latihan yang tercatat untuk hari %s\n", hari)
	}
}

func pilihHariInput() string {
	fmt.Println("1. Senin")
	fmt.Println("2. Selasa")
	fmt.Println("3. Rabu")
	fmt.Println("4. Kamis")
	fmt.Println("5. Jumat")
	fmt.Println("6. Sabtu")
	fmt.Println("7. Minggu")

	var pilihan int
	fmt.Print("Pilih hari (1-7): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		return "senin"
	case 2:
		return "selasa"
	case 3:
		return "rabu"
	case 4:
		return "kamis"
	case 5:
		return "jumat"
	case 6:
		return "sabtu"
	case 7:
		return "minggu"
	default:
		fmt.Println("Pilihan tidak valid, default ke Senin")
		return "senin"
	}
}

func tambahTracking(A *tabWorkout, n int) {
	fmt.Println("\n=== Tambah Data Tracking ===")
	for i := 0; i < n; i++ {
		fmt.Printf("\nUntuk %s\n", A[i].nama)

		A[i].hari = pilihHariInput()
		fmt.Print("Masukkan jumlah kalori harian: ")
		fmt.Scan(&A[i].kalori)

		fmt.Print("Masukkan jumlah protein harian (g): ")
		fmt.Scan(&A[i].protein)

		fmt.Print("Masukkan jam tidur: ")
		fmt.Scan(&A[i].jamtidur)

		fmt.Println("Data tracking berhasil ditambahkan!")
	}
}

func tampilkanTracking(A tabWorkout, n int) {
	fmt.Println("\n=== Data Tracking ===")
	for i := 0; i < n; i++ {
		if A[i].kalori != 0 || A[i].protein != 0 {
			fmt.Printf("\n%s\n", A[i].nama)
			fmt.Printf("Kalori: %d kkal\n", A[i].kalori)
			fmt.Printf("Protein: %d g\n", A[i].protein)
			fmt.Printf("Jam tidur: %.1f jam\n", A[i].jamtidur)
		}
	}
}

func pilihHari(A tabWorkout, n int) {
	fmt.Println("\n=== Tampilkan Data per Hari ===")
	hari := pilihHariInput()

	fmt.Printf("\n=== Data Hari %s ===\n", hari)
	found := false

	for i := 0; i < n; i++ {
		if A[i].hari == hari {
			found = true
			fmt.Printf("\n%s\n", A[i].nama)

			// Tampilkan profil
			fmt.Println("\nProfil:")
			fmt.Printf("Gender: %s\n", A[i].gender)
			fmt.Printf("Tujuan: %s\n", A[i].tujuan)
			fmt.Printf("Tinggi/Berat: %.1f cm / %.1f kg\n", A[i].tinggi, A[i].berat)

			// Tampilkan latihan jika ada
			if A[i].latihan != "" {
				fmt.Println("\nLatihan:")
				fmt.Printf("Jenis: %s\n", A[i].latihan)
				fmt.Printf("Rep: %d, Set: %d\n", A[i].rep, A[i].set)
				fmt.Printf("Beban: %.1f kg\n", A[i].beban)
			}

			// Tampilkan tracking jika ada
			if A[i].kalori != 0 || A[i].protein != 0 {
				fmt.Println("\nTracking:")
				fmt.Printf("Kalori: %d kkal\n", A[i].kalori)
				fmt.Printf("Protein: %d g\n", A[i].protein)
				fmt.Printf("Jam tidur: %.1f jam\n", A[i].jamtidur)
			}
		}
	}

	if !found {
		fmt.Printf("Tidak ada data untuk hari %s\n", hari)
	}
}

func urutkanHariByBebanDescending(A tabWorkout, n int) {
	fmt.Println("\n=== Urutan Hari dari Latihan Terberat ke Teringan ===")

	// Buat array untuk menyimpan hari dan beban maksimumnya
	type hariBeban struct {
		hari  string
		beban float64
	}
	var hariBebanList [7]hariBeban

	// Inisialisasi hari
	hariBebanList[0].hari = "senin"
	hariBebanList[1].hari = "selasa"
	hariBebanList[2].hari = "rabu"
	hariBebanList[3].hari = "kamis"
	hariBebanList[4].hari = "jumat"
	hariBebanList[5].hari = "sabtu"
	hariBebanList[6].hari = "minggu"

	// Hitung beban maksimum per hari
	for i := 0; i < n; i++ {
		for j := 0; j < 7; j++ {
			if A[i].hari == hariBebanList[j].hari {
				if A[i].beban > hariBebanList[j].beban {
					hariBebanList[j].beban = A[i].beban
				}
				break
			}
		}
	}

	// Urutkan dari terbesar ke terkecil menggunakan insertion sort
	for i := 1; i < 6; i++ {
		key := hariBebanList[i]
		j := i - 1
	
		for j >= 0 && hariBebanList[j].beban < key.beban {
			hariBebanList[j+1] = hariBebanList[j]
			j--
		}
	
		hariBebanList[j+1] = key
	}

	//Tampilkan hasil
	for i := 0; i < 7; i++ {
		if hariBebanList[i].beban > 0 {
			fmt.Printf("%d. Hari %s - Beban Maksimum: %.1f kg\n", i+1, hariBebanList[i].hari, hariBebanList[i].beban)
		} else {
			fmt.Printf("%d. Hari %s - Tidak ada data latihan\n", i+1, hariBebanList[i].hari)
		}
	}
}
func urutkanHariByBebanAscending(A tabWorkout, n int) {
	fmt.Println("\n=== Urutan Hari dari Latihan Teringan ke Terberat ===")

	// Buat array untuk menyimpan hari dan beban maksimumnya
	type hariBeban struct {
		hari  string
		beban float64
	}
	var hariBebanList [7]hariBeban

	// Inisialisasi hari
	hariBebanList[0].hari = "senin"
	hariBebanList[1].hari = "selasa"
	hariBebanList[2].hari = "rabu"
	hariBebanList[3].hari = "kamis"
	hariBebanList[4].hari = "jumat"
	hariBebanList[5].hari = "sabtu"
	hariBebanList[6].hari = "minggu"

	// Hitung beban maksimum per hari
	for i := 0; i < n; i++ {
		for j := 0; j < 7; j++ {
			if A[i].hari == hariBebanList[j].hari {
				if A[i].beban > hariBebanList[j].beban {
					hariBebanList[j].beban = A[i].beban
				}
				break
			}
		}
	}

	// Urutkan dari terkecil ke terbesar menggunakan insertion sort
	for i := 1; i < 6; i++ {
		key := hariBebanList[i]
		j := i - 1
	
		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && hariBebanList[j].beban > key.beban {
			hariBebanList[j+1] = hariBebanList[j]
			j--
		}
	
		hariBebanList[j+1] = key
	}

	// Tampilkan hasil
	for i := 0; i < 7; i++ {
		if hariBebanList[i].beban > 0 {
			fmt.Printf("%d. Hari %s - Beban Maksimum: %.1f kg\n", i+1, hariBebanList[i].hari, hariBebanList[i].beban)
		} else {
			fmt.Printf("%d. Hari %s - Tidak ada data latihan\n", i+1, hariBebanList[i].hari)
		}
	}
}
