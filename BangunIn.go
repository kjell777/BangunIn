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
	//variabel untuk menyimpan banyaknya array rpelayanan yang terisi pada setiap supplier, sehingga ketika output riwayat pelayanan tidak perlu mengecek indeks 1 - NMAX untuk output indeks yang memiliki isi
	//riwayatke baru ditambah setelah input rpelayanan ke 2/indeks 1 dan seterusnya atau setelah fungsi inputrpelayanan
	//lebih hemat memori dan waktu karena tidak perlu mengecek indeks yang tidak terisi pada array rpelayanan, cukup print indeks 1 - riwayatke saja (indeks 0 ikut di print di awal dengan data suppliers lain)
}
type wilayah struct {
	nama        string
	isiSupplier supplier
	wilcount    int //banyaknya supplier yang termasuk ke dalam 1 wilayah
}
type wilayahs [NMAX]wilayah
type supplier [NMAX]suppliers

//wilayah -> nama : Jakarta
//			isiSupplier :

func main() {
	var s supplier
	var arrCount int
	menuawal(s, arrCount) // program dimulai dari menuawal
}

//  menu awal aplikasi
func menuawal(s supplier, arrCount int) {
	var pilihan int
	fmt.Printf("\n=================================================\n")
	fmt.Printf("                    BANGUNIN                   \n")
	fmt.Printf("=================================================\n")
	fmt.Printf("[1] Open List Supplier\n")
	fmt.Printf("[2] Statistik Supplier\n")
	fmt.Printf("[3] Quit\n")
	fmt.Printf("=================================================\n")
	fmt.Printf("Pilih menu (1-3): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		menuutama(s, arrCount)
	case 2:
		countWilayah(s, arrCount)
	case 3:
		fmt.Printf("\nProgram selesai.\n")
	default:
		fmt.Printf("\nInputan tidak valid.\n")
		menuawal(s, arrCount)
	}
}

//menu modifikasi dan ganti" supplier
func menuutama(s supplier, arrCount int) {
	outputDaftarMitra(s, arrCount)
	var pilihan int
	fmt.Printf("\n=====================================================================================================\n")
	fmt.Printf("MENU : \n")
	fmt.Printf("[1] Tambah Data Supplier (ADD)\n")
	fmt.Printf("[2] Modifikasi Data (MODIFICATION / DELETE)\n")
	fmt.Printf("[3] Urutkan Data (SORT)\n")
	fmt.Printf("[4] Cari Data (SEARCH)\n")
	fmt.Printf("[5] Kembali ke Menu Awal (RETURN)\n")
	fmt.Printf("Pilih aksi (1-5): ")
	fmt.Scan(&pilihan)
	if pilihan >= 2 && pilihan <= 4 && arrCount == 0 {
		fmt.Printf("\nData masih kosong, silakan input data terlebih dahulu.\n")
		menuutama(s, arrCount)
	}
	switch pilihan {
	case 1:
		inputData(&s, &arrCount)
		menuutama(s, arrCount)
	case 2:
		modification(&s, &arrCount)
		menuutama(s, arrCount)
	case 3:
		selectSorting(&s, arrCount)
		menuutama(s, arrCount)
	case 4:
		selectSearch(s, arrCount)
		kembalikemenuutama(s, arrCount)
	case 5:
		menuawal(s, arrCount)
	default:
		fmt.Printf("\nInputan tidak valid, silakan coba lagi.\n")
		menuutama(s, arrCount)
	}
}

func kembalikemenuutama(s supplier, arrCount int) {
	var kembali string
	fmt.Printf("\n[1] RETURN")
	fmt.Scan(&kembali)
	menuutama(s, arrCount)
}

func inputData(s *supplier, arrCount *int) {
	//arrCount dibuat (*) untuk dapat terus mengupdate banyaknya array supplier yang terisi, karena array supplier dapat terus dimodifikasi (add, delete)
	fmt.Print("\nPETUNJUK : GUNAKAN UNDERSCORE (_) SEBAGAI SPASI")
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
	//sebagai fungsi lanjutan untuk akses opsi pada outputDaftarMitra
	var i int
	fmt.Printf("\n%2s |       %4s      |    %5s    |      %6s     |     %8s    | %5s | %s\n-----------------------------------------------------------------------------------------------------\n", "NO", "NAMA", "KONTAK", "LOKASI", "MATERIAL", "RATING", "RIWAYAT PELAYANAN")
	if s[i].nama != "" {
		for i := 0; i < arrCount; i++ {
			fmt.Printf("%-2d | %-15s | %-12s | %-15s | %-15s |  %-.2f  | %s\n", i+1, s[i].nama, s[i].kontak, s[i].wilayah, s[i].jenismat, s[i].rating, s[i].rpelayanan[0])
			if s[i].riwayatke > 0 {
				for j := 1; j <= s[i].riwayatke; j++ {
					fmt.Printf("%-2s | %-15s | %-12s | %-15s | %-15s |  %-3s   | %s\n", " ", " ", " ", " ", " ", " ", s[i].rpelayanan[j])
				}
			}
			fmt.Printf("-----------------------------------------------------------------------------------------------------\n")
		}
	} else {
		fmt.Printf("\n%32sBELUM ADA DATA YANG DIINPUT !!!%33s\n", "", "")
	}
}

func modification(s *supplier, arrCount *int) {
	//fungsi untuk modifikasi data supplier, s dan arrcount dibuat pointer (*) karena modifikasi akan mengubah jumlah/isi dari array s dan juga arrCount yang menyimpan banyaknya array supplier yang terisi jadi harus terus diupdate pada variabel dimana fungsi ini akan dipanggil (outputDaftarMitra)
	var replace string
	var num, num2, choice, category int
	var r3place float64
	fmt.Printf("\n[1] MODIFIKASI\n[2] DELETE\nPilih aksi (1-2): ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("\n[1] NAMA\n[2] KONTAK\n[3] JENIS MATERIAL\n[4] WILAYAH\n[5] RIWAYAT PELAYANAN\n[6] RATING\nVariabel yang ingin di modifikasi (tulis dalam ANGKA): ")
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
			fmt.Printf("\n[1] MODIF RIWAYAT\n[2] ADD RIWAYAT\nPilih aksi (1-2): ")
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
	//fungsi untuk modifikasi riwayat pelayanan, num dan num2 sebagai indeks supplier dan riwayat pelayanan yang ingin dimodifikasi
	//s dibuat pointer (*) karena modifikasi akan mengubah isi dari array s, jadi harus terus diupdate pada variabel dimana fungsi ini akan dipanggil (modification)
	var replace string
	fmt.Printf("\nMasukkan riwayat pelayanan baru: ")
	fmt.Scan(&replace)
	s[num].rpelayanan[num2-1] = replace
}

//variabel s pada semua fungsi sorting dibuat pointer (*) karena sorting akan mengubah urutan isi dari array s, jadi harus terus diupdate pada variabel dimana fungsi ini akan dipanggil
func selectSorting(s *supplier, arrCount int) {
	//untuk memilih sorting asc or desc (by rating, untuk opsi by variabel lain menyusul)
	var option int
	fmt.Printf("\n[1] ASCENDING\n[2] DESCENDING\nPilih opsi (1-2): ")
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

//variabel pada fungsi search tidak perlu dibuat pointer dikarenbakan fungsi search tidak akan mengubah isi dari array supplier (hanya mencari dan menampilkan yang dicari)
func selectSearch(s supplier, arrCount int) {
	//fungsi untuk memilih SEARCH BY(name or lokasi...mungkin akan ditambahkan opsi lain setelah review)
	fmt.Printf("\n[1] SEARCH BY NAME\n[2] SEARCH BY LOCATION\nPilih opsi (1-2): ")
	var find string
	var option int
	fmt.Scan(&option)
	fmt.Printf("\nMasukkan keyword pencarian: ")
	fmt.Scan(&find)
	switch option {
	case 1:
		BinarySearchNama(s, arrCount, find)
	case 2:
		SequentialSearchLokasi(s, arrCount, find)
	}
	// [HAPUS] input return manual diganti
}

func outputSearch(sSearch supplier, searchcount int) {
	//output hasil search, untuk memimasahkan tabel utama dengan hasil search (Agar bisa kembali ke tabel utama setelah melihat hasil search)
	fmt.Printf("\n%2s |       %4s      |    %5s    |     %6s      |     %8s    | %5s | %s\n-------------------------------------------------------------------------------------------------\n", "NO", "NAMA", "KONTAK", "LOKASI", "MATERIAL", "RATING", "RIWAYAT PELAYANAN")
	var i int
	if sSearch[i].nama != "" {
		for i := 0; i < searchcount; i++ {
			fmt.Printf("%-2d | %-15s | %-12s | %-15s | %-15s |  %-.2f  | %s\n", i+1, sSearch[i].nama, sSearch[i].kontak, sSearch[i].wilayah, sSearch[i].jenismat, sSearch[i].rating, sSearch[i].rpelayanan[0])
			if sSearch[i].riwayatke > 0 {
				for j := 1; j <= sSearch[i].riwayatke; j++ {
					fmt.Printf("%-2s | %-15s | %-12s | %-15s | %-15s |  %-3s  | %s\n", " ", " ", " ", " ", " ", " ", sSearch[i].rpelayanan[j])
				}
			}
			fmt.Println("-------------------------------------------------------------------------------------------------")
			// [BARU] bikin "sekat' perbarisnya
		}
	} else {
		fmt.Printf("\n%16sDATA TIDAK DITEMUKAN !!!%17s\n", "", "")
	}
}

func SequentialSearchLokasi(s supplier, arrCount int, find string) {
	//mencari supplier via lokasi menggunakan sequential search, memasukkan supplier yang memiliki lokasi yg sama ke array lokasi lalu output
	var lokasi supplier
	var arrLokasi int
	for i := 0; i < arrCount; i++ {
		if s[i].wilayah == find {
			lokasi[arrLokasi] = s[i] //dilakukan untuk memindahkan isi dari supplier dicari (via lokasi) ke array lokasi pada indeks ke arrLokasi (supplier yang berbeda bisa mempunyai lokasi yang sama, maka dibuat array baru agar pada proses output bisa keluar semua supplier dengan lokasi yang dicari)

			arrLokasi++ //menambah jumlah data yang berhasil ditemukan sekaligus menjadi indeks berikutnya untuk array lokasi
		}
	}
	outputSearch(lokasi, arrLokasi)
}

func BinarySearchNama(s supplier, arrCount int, find string) {
	//mencari supplier berdasarkan nama, disort by nama terlebih dahulu
	SelectionSortNama(&s, arrCount)
	//nilai array supplier pada fungsi selectionsortnama akan dicopy ke variabel lokal dari binarysearch (s) sehingga tidak mengubah isi dari array asli supplier pada outputDaftarMitra/tabel supplier asli)

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

func countWilayah(s supplier, arrCount int) {
	//menghitung banyaknya wilayah yang ada (countw) dan masukkan nama wilayah ke dalam array wilayah.nama (tanpa dupe/double)
	var w wilayahs
	var tanda bool
	var countw int
	var returns string
	w[0].nama = s[0].wilayah
	countw = 0
	for i := 1; i < arrCount; i++ {
		tanda = false
		for temp := 0; temp <= countw; temp++ {
			if w[temp].nama == s[i].wilayah {
				tanda = true
			}
		}
		if !tanda {
			countw++
			w[countw].nama = s[i].wilayah
		}
	}
	//untuk mengisi array setiap wilayah (w[0-*countw].isiSupplier[0-w[i].wilcount-1]) dengan nilai array supplier dengan wilayah yang sama
	for i := 0; i <= countw; i++ {
		for j := 0; j < arrCount; j++ {
			if w[i].nama == s[j].wilayah {
				w[i].isiSupplier[w[i].wilcount] = s[j]
				w[i].wilcount++
			}
		}
	}
	outputWilayah(w, countw)
	fmt.Print("\n[1] RETURN\n")
	fmt.Scan(&returns)
	menuawal(s, arrCount)
}

func outputWilayah(w wilayahs, countw int) {
	//output jumlah dan rata2 wilayah]
	if w[0].nama == "" {
		fmt.Printf("\n%16sBELUM ADA DATA YANG DIINPUT !!!%17s\n", "", "")
	} else {
		for i := 0; i <= countw; i++ {
			fmt.Printf("\nWilayah %s: jumlah %d supplier dan rata-rata rating adalah %.2f", w[i].nama, w[i].wilcount, rata2rating(w, i))
			fmt.Printf("\n%2s |       %4s      |    %5s    |      %6s     |     %8s    | %5s | %s\n-------------------------------------------------------------------------------------------------\n", "NO", "NAMA", "KONTAK", "LOKASI", "MATERIAL", "RATING", "RIWAYAT PELAYANAN")
			for j := 0; j < w[i].wilcount; j++ {
				fmt.Printf("%-2d | %-15s | %-12s | %-15s | %-15s |  %-.2f  | %s\n", j+1, w[i].isiSupplier[j].nama, w[i].isiSupplier[j].kontak, w[i].isiSupplier[j].wilayah, w[i].isiSupplier[j].jenismat, w[i].isiSupplier[j].rating, w[i].isiSupplier[j].rpelayanan[0])
				if w[i].isiSupplier[j].riwayatke > 0 {
					for k := 1; k <= w[i].isiSupplier[j].riwayatke; k++ {
						fmt.Printf("%-2s | %-15s | %-12s | %-15s | %-15s |  %-4s  | %s\n", " ", " ", " ", " ", " ", " ", w[i].isiSupplier[j].rpelayanan[k])
					}
				}
			}
		}
	}
}

func rata2rating(w wilayahs, i int) float64 {
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

//catatan asistensi : menu dibuat didepan (jadi ga langsung output tabel), menu dibuat kebawah (karena kaatanya kalau kesamping orang bakal ga ngeh) dan bikin "sekat" per baris outputnya
