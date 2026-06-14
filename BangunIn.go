package main

import "fmt"

const NMAX int = 100

type suppliers struct {
	nama, kontak      string
	rating            float64
	jenismat, wilayah string
	rpelayanan        [NMAX]string
	//array untuk menyiman riwayat pelayanan dari setiap supplier (banyaknya riwayat pelayanan yg dipunyai supplier disimpan pada array riwayatke dengan indeks yang sesuai dengan supplier pada array supplier, contoh : supplier[0] = supplier A, riwayatke[0] = banyaknya riwayat pelayanan yang dimiliki supplier A)
	//1 supplier dapat memiliki > 1 riwayat pelayanan karena mungkin mitra dari pemakai aplikasi bekerja sama dengan supplier yang sama > 1 kali dan/atau memakai jasa pelayanan yang berbeda dari supplier yang sama di waktu yang berbeda/sama.
	riwayatke int
	//array untuk menyimpan banyaknya array rpelayanan yang terisi pada setiap supplier, sehingga ketika output riwayat pelayanan tidak perlu mengecek indeks 1 - NMAX untuk output indeks yang memiliki isi
	//lebih hemat memori dan waktu karena tidak perlu mengecek indeks yang tidak terisi pada array rpelayanan, cukup print indeks 1 - riwayatke saja (indeks 0 ikut di print di awal dengan data suppliers lain)
}
type wilayah struct {
	nama        string
	isiSupplier supplier
	wilcount    int //banyaknya supplier yang termasuk ke dalam 1 wilayah
}
type wilayahs [NMAX]wilayah
type supplier [NMAX]suppliers

func main() {
	var s supplier
	var arrCount int
	outputDaftarMitra(s, arrCount)
}
func inputData(s *supplier, arrCount *int) {
	//arrCount dibuat (*) untuk dapat terus mengupdate banyaknya array supplier yang terisi, karena array supplier dapat terus dimodifikasi (add, delete)
	fmt.Print("PETUNJUK : GUNAKAN UNDERSCORE (_) SEBAGAI SPASI")
	fmt.Printf("\nMasukkan nama supplier: ")
	fmt.Scan(&s[*arrCount].nama)
	fmt.Printf("\nMasukkan kontak supplier: ")
	fmt.Scan(&s[*arrCount].kontak)
	fmt.Printf("\nMasukkan wilayah supplier: ")
	fmt.Scan(&s[*arrCount].wilayah)
	fmt.Printf("\nMasukkan jenis material yang disupply: ")
	fmt.Scan(&s[*arrCount].jenismat)
	fmt.Printf("\nMasukkan rating supplier (0.0 - 5.0): ")
	fmt.Scan(&s[*arrCount].rating)
	for s[*arrCount].rating < 0 || s[*arrCount].rating > 5 {
		fmt.Printf("\nRating tidak valid, silakan masukkan rating kembali.")
		fmt.Printf("\nMasukkan rating supplier (0.0 - 5.0): ")
		fmt.Scan(&s[*arrCount].rating)
	}
	fmt.Printf("\nMasukkan riwayat pelayanan supplier: ")
	fmt.Scan(&s[*arrCount].rpelayanan[0])
	*arrCount++
}
func inputRPelayanan(s *supplier, num int) {
	//riwayatke (*) berfungsi untuk mengupdate banyaknya riwayat pelayanan yang ditambahkan pada satu supplier
	s[num].riwayatke = s[num].riwayatke + 1
	fmt.Printf("\nMasukkan riwayat pelayanan supplier: ")
	fmt.Scan(&s[num].rpelayanan[s[num].riwayatke])
}
func outputDaftarMitra(s supplier, arrCount int) {
	fmt.Printf("\n%2s |       %4s      |    %5s    |   %6s   |     %8s    | %5s | %s\n", "NO", "NAMA", "KONTAK", "LOKASI", "MATERIAL", "RATING", "RIWAYAT PELAYANAN")
	tanda := false
	for i := 0; i < arrCount; i++ {
		if s[i].nama != "" {
			fmt.Printf("%-2d | %-15s | %-12s | %-10s | %-15s |  %-.2f  | %s\n", i+1, s[i].nama, s[i].kontak, s[i].wilayah, s[i].jenismat, s[i].rating, s[i].rpelayanan[0])
			if s[i].riwayatke > 0 {
				for j := 1; j <= s[i].riwayatke; j++ {
					fmt.Printf("%-2s | %-15s | %-12s | %-10s | %-15s |  %-3s   | %s\n", " ", " ", " ", " ", " ", " ", s[i].rpelayanan[j])
				}
			}
			tanda = true
		}
	}
	if !tanda {
		fmt.Printf("\n%16sBELUM ADA DATA YANG DIINPUT !!!%17s", "", "")
	}
	fmt.Printf("\n\nADD Supplier [1] | MODIFICATION [2] | SORT [3] | SEARCH [4] | STATISTICS [5] | QUIT [6]\n")
	DaftarMitraLanjutan(s, arrCount)
}
func DaftarMitraLanjutan(s supplier, arrCount int) {
	//sebagai fungsi lanjutan untuk akses opsi pada outputDaftarMitra
	var lanjutan int
	var w wilayahs
	var num int
	fmt.Scan(&lanjutan)
	switch lanjutan {
	case 1:
		inputData(&s, &arrCount)
		outputDaftarMitra(s, arrCount)
	case 2:
		modification(&s, &arrCount)
		outputDaftarMitra(s, arrCount)
	case 3:
		selectSorting(&s, arrCount)
		outputDaftarMitra(s, arrCount)
	case 4:
		selectSearch(s, arrCount)
	case 5:
		countWilayah(s, &w, arrCount, &num)
	case 6:
		fmt.Printf("\nProgram selesai.")
	default:
		fmt.Printf("\nInputan tidak valid, silakan coba lagi.")
		DaftarMitraLanjutan(s, arrCount)
	}
}
func modification(s *supplier, arrCount *int) {
	var replace string
	var num, num2, choice, category int
	var r3place float64
	fmt.Printf("\nMODIFICATION [1] | DELETE [2] : ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("\nNAMA [1] | KONTAK [2] | JENIS MATERIAL [3] | WILAYAH [4] | RIWAYAT PELAYANAN [5] | RATING [6]\nVariabel yang ingin di modifikasi (tulis dalam ANGKA): ")
		fmt.Scan(&category)
		fmt.Printf("\nMasukkan nomor supplier yang ingin di modifikasi: ")
		fmt.Scan(&num)
		for num < 1 || num > *arrCount {
			fmt.Printf("\nNomor supplier tidak valid, silakan coba lagi.")
			fmt.Printf("\nMasukkan nomor supplier yang ingin di modifikasi: ")
			fmt.Scan(&num)
		}
		num = num - 1
		if category >= 1 && category <= 4 {
			fmt.Printf("\nMasukkan isi baru: ")
			fmt.Scan(&replace)
			switch category {
			case 1:
				s[num].nama = replace
			case 2:
				s[num].kontak = replace
			case 3:
				s[num].jenismat = replace
			case 4:
				s[num].wilayah = replace
			}
		} else if category == 5 {
			fmt.Printf("\nMODIF RIWAYAT [1] | ADD RIWAYAT [2]: ")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Printf("\nMasukkan isi baru: ")
				fmt.Scan(&replace)
				if s[num].riwayatke == 0 {
					s[num].rpelayanan[0] = replace
				} else {
					fmt.Printf("\nMasukkan nomor riwayat pelayanan yang ingin di modifikasi (1 - %d): ", s[num].riwayatke)
					fmt.Scan(&num2)
					for num2 < 1 || num2 > s[num].riwayatke {
						fmt.Printf("\nNomor tidak valid, coba lagi: ")
						fmt.Scan(&num2)
					}
					modificationRP(s, num, num2)
				}
			case 2:
				inputRPelayanan(s, num)
			}
		} else if category == 6 {
			fmt.Printf("\nMasukkan rating baru: ")
			fmt.Scan(&r3place)
			for r3place < 0 || r3place > 5 {
				fmt.Printf("\nRating tidak valid, silakan masukkan rating kembali.")
				fmt.Printf("\nMasukkan rating baru: ")
				fmt.Scan(&r3place)
			}
			s[num].rating = r3place
		} else {
			fmt.Print("KATEGORI TIDAK DITEMUKAN !! ")
		}
	case 2:
		var idxdel int
		fmt.Printf("\nSupplier yang ingin di DELETE (tulis nomor): ")
		fmt.Scan(&idxdel)
		for idxdel < 1 || idxdel > *arrCount {
			fmt.Printf("\nNomor supplier tidak valid, silakan coba lagi.")
			fmt.Printf("\nSupplier yang ingin di DELETE (tulis nomor): ")
			fmt.Scan(&idxdel)
		}
		idxdel = idxdel - 1
		for i := idxdel; i < *arrCount-1; i++ {
			s[i] = s[i+1] //menggeser supplier ke kiri untuk mendelete supplier yang ingin didelete
		}
		s[*arrCount-1] = suppliers{}
		//ngosongin array supplier terakhir,  s[*arrCount-1] <= indeks satuan, tipe datanya struct suppliers makanya pake suppliers{}
		*arrCount--
	}
}
func modificationRP(s *supplier, num int, num2 int) {
	var replace string
	fmt.Printf("\nMasukkan riwayat pelayanan baru: ")
	fmt.Scan(&replace)
	s[num].rpelayanan[num2-1] = replace
}
func selectSorting(s *supplier, arrCount int) {
	//untuk memilih sorting asc or desc (by rating, untuk opsi by variabel lain menyusul)
	var option int
	fmt.Printf("\nASCENDING [1] || DESCENDING [2] by rating (tulis dalam ANGKA):")
	fmt.Scan(&option)
	switch option {
	case 1:
		SelectionSortAsc(s, arrCount)
	case 2:
		InsertionSortDesc(s, arrCount)
	}
}
func SelectionSortAsc(s *supplier, arrCount int) {
	// selection sort ascending based on rating
	var idx int
	for i := 0; i < arrCount; i++ {
		idx = i
		for j := i + 1; j < arrCount; j++ {
			if s[idx].rating > s[j].rating {
				idx = j
			}
		}
		temp := s[i]
		s[i] = s[idx]
		s[idx] = temp
		// menggunakam temp agar tidak ada duplikat data
	}
}
func InsertionSortDesc(s *supplier, arrCount int) {
	// insertion sort descending based on rating
	for pass := 1; pass <= arrCount-1; pass++ {
		i := pass
		temp := s[i]
		for i > 0 && temp.rating > s[i-1].rating {
			s[i] = s[i-1]
			i--
		}
		s[i] = temp
	}
}
func SelectionSortNama(s *supplier, arrCount int) {
	//selection sort ascending based on nama (untuk kebutuhan BinarySearchNama)
	for i := 0; i < arrCount; i++ {
		idx := i
		for j := i + 1; j < arrCount; j++ {
			if s[idx].nama > s[j].nama {
				idx = j
			}
		}
		temp := s[i]
		s[i] = s[idx]
		s[idx] = temp
		// menggunakam temp agar tidak ada duplikat data
	}
}
func selectSearch(s supplier, arrCount int) {
	//fungsi untuk memilih SEARCH BY(name or lokasi...mungkin akan ditambahkan opsi lain setelah review)
	fmt.Printf("\nSEARCH BY NAME [1] | LOCATION [2] :")
	var find string
	var option, pilihan int
	fmt.Scan(&option)
	fmt.Printf("\nMasukkan keyword pencarian: ")
	fmt.Scan(&find)
	switch option {
	case 1:
		BinarySearchNama(s, arrCount, find)
	case 2:
		SequentialSearchLokasi(s, arrCount, find)
	}
	fmt.Printf("\n\nRETURN [1]\n")
	fmt.Scan(&pilihan)
	outputDaftarMitra(s, arrCount)
}
func outputSearch(sSearch supplier, searchcount int) {
	//output hasil search, untuk memimasahkan tabel utama dengan hasil search (Agar bisa kembali ke tabel utama setelah melihat hasil search)
	fmt.Printf("\n%2s |       %4s      |    %5s    |   %6s   |     %8s    | %5s | %s\n", "NO", "NAMA", "KONTAK", "LOKASI", "MATERIAL", "RATING", "RIWAYAT PELAYANAN")
	tanda := false
	for i := 0; i < searchcount; i++ {
		if sSearch[i].nama != "" {
			fmt.Printf("%-2d | %-15s | %-12s | %-10s | %-15s |  %-.2f  | %s\n", i+1, sSearch[i].nama, sSearch[i].kontak, sSearch[i].wilayah, sSearch[i].jenismat, sSearch[i].rating, sSearch[i].rpelayanan[0])
			if sSearch[i].riwayatke > 0 {
				for j := 1; j <= sSearch[i].riwayatke; j++ {
					fmt.Printf("%-2s | %-15s | %-12s | %-10s | %-15s |  %-3s  | %s\n", " ", " ", " ", " ", " ", " ", sSearch[i].rpelayanan[j])
				}
			}
			tanda = true
		}
	}
	if !tanda {
		fmt.Printf("\n%16sDATA TIDAK DITEMUKAN !!!%17s", "", "")
	}
}
func SequentialSearchLokasi(s supplier, arrCount int, find string) {
	//mencari supplier via lokasi menggunakan sequential search, memasukkan supplier yang memiliki lokasi yg sama ke array lokasi lalu output
	var lokasi supplier
	var arrLokasi int
	for i := 0; i < arrCount; i++ {
		if s[i].wilayah == find {
			lokasi[arrLokasi] = s[i] //dilakukan untuk memindahkan isi dari supplier dicari (via lokasi) ke array lokasi pada indeks ke arrLokasi (supplier yang berbeda bisa mempunyai lokasi yang sama, maka dibuat array baru agar pada proses output bisa keluar semua supplier dengan lokasi yang dicari)
			arrLokasi++              //menambah jumlah data yang berhasil ditemukan sekaligus menjadi indeks berikutnya untuk array lokasi
		}
	}
	outputSearch(lokasi, arrLokasi)
}
func BinarySearchNama(s supplier, arrCount int, find string) {
	//mencari supplier berdasarkan nama, disort by nama terlebih dahulu
	SelectionSortNama(&s, arrCount)
	var nama supplier
	var left, mid, right, found, arrNama, i int
	found = -1
	left = 0
	right = arrCount - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if s[mid].nama == find {
			found = mid
		} else if s[mid].nama > find {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if found != -1 {
		nama[arrNama] = s[found]
		arrNama++
		i = mid - 1
		for i >= 0 && s[i].nama == find {
			nama[arrNama] = s[i]
			arrNama++
			i--
		}
		i = mid + 1
		for i < arrCount && s[i].nama == find {
			nama[arrNama] = s[i]
			arrNama++
			i++
		}
	}
	outputSearch(nama, arrNama)
}
func countWilayah(s supplier, w *wilayahs, arrCount int, countw *int) {
	//menghitung banyaknya wilayah yang ada dan masukkan ke dalam array wilayah.nama (tanpa dupe/double)
	var tanda bool
	var pilihan int
	w[0].nama = s[0].wilayah
	*countw = 0
	for i := 1; i < arrCount; i++ {
		tanda = false
		for temp := 0; temp <= *countw; temp++ {
			if w[temp].nama == s[i].wilayah {
				tanda = true
			}
		}
		if !tanda {
			*countw++
			w[*countw].nama = s[i].wilayah
		}
	}
	//untuk mengisi array setiap wilayah dengan nilai array supplier dengan wilayah yang sama (Contoh : W[0].nama = Jabar, isiSupplier [0] = array supplier abcd, isiSupplier [1] = array supplier abc && W[1].nama = Jakarta, isiSupplier [0]  = supplier cde
	for i := 0; i <= *countw; i++ {
		for j := 0; j < arrCount; j++ {
			if w[i].nama == s[j].wilayah {
				w[i].isiSupplier[w[i].wilcount] = s[j]
				w[i].wilcount++
			}
		}
	}
	outputWilayah(w, *countw)
	fmt.Printf("\n\nRETURN [1]\n")
	fmt.Scan(&pilihan)
	outputDaftarMitra(s, arrCount)
}
func outputWilayah(w *wilayahs, countw int) {
	//output jumlah dan rata2 wilayah
	for i := 0; i <= countw; i++ {
		fmt.Printf("WIlayah %s: jumlah %d supplier dan rata-rata rating adalah %.2f\n", w[i].nama, w[i].wilcount, rata2rating(w, i))
		for j := 0; j < w[i].wilcount; j++ {
			fmt.Printf("%-2d | %-15s | %-12s | %-10s | %-15s |  %-.2f  | %s\n", j+1, w[i].isiSupplier[j].nama, w[i].isiSupplier[j].kontak, w[i].isiSupplier[j].wilayah, w[i].isiSupplier[j].jenismat, w[i].isiSupplier[j].rating, w[i].isiSupplier[j].rpelayanan[0])
			if w[i].isiSupplier[j].riwayatke > 0 {
				for k := 1; k <= w[i].isiSupplier[j].riwayatke; k++ {
					fmt.Printf("%-2s | %-15s | %-12s | %-10s | %-15s |  %-3s  | %s\n", " ", " ", " ", " ", " ", " ", w[i].isiSupplier[j].rpelayanan[k])
				}
			}
		}
	}
}
func rata2rating(w *wilayahs, i int) float64 {
	//rata-rata rating pada wilayah ke i
	var rata2 float64
	for j := 0; j < w[i].wilcount; j++ {
		rata2 = rata2 + w[i].isiSupplier[j].rating
	}
	if w[i].wilcount == 0 {
		return 0
	}
	return rata2 / float64(w[i].wilcount)
}

//catatan asistensi : menu dibuat didepan (jadi ga langsung output tabel), menu dibuat kebawah (karena kaatanya kalau kesamping orang bakal ga ngeh)
