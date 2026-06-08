package main
import "fmt"
const NMAX int = 100
type suppliers struct {
	nama, kontak      string
	rating            float64
	jenismat, wilayah string
	rpelayanan        [NMAX]string
}
type wilayah struct {
	nama        string
	isiSupplier supplier
	wilcount    int //batas array yang terisi (mulai dari 0)
	riwayatpw   [NMAX]int
}
type wilayahs [NMAX]wilayah
type supplier [NMAX]suppliers
func main() {
	var s supplier
	var arrCount int
	var riwayatke [NMAX]int
	outputDaftarMitra(s, arrCount, riwayatke)
}
func inputData(s *supplier, arrCount *int) {
	//arrCount dibuat (*) untuk dapat terus mengupdate banyaknya array supplier yang terisi, karena array supplier dapat terus dimodifikasi (add, delete)
	fmt.Scan(&s[*arrCount].nama, &s[*arrCount].kontak, &s[*arrCount].wilayah, &s[*arrCount].jenismat, &s[*arrCount].rating, &s[*arrCount].rpelayanan[0])
	*arrCount++
}
func inputRPelayanan(s *supplier, num int, riwayatke *[NMAX]int) {
	//riwayatke (*) berfungsi untuk mengupdate banyaknya riwayat pelayanan yang ditambahkan pada satu supplier
	riwayatke[num] = riwayatke[num] + 1
	fmt.Scan(&s[num].rpelayanan[riwayatke[num]])
}
func outputDaftarMitra(s supplier, arrCount int, riwayatke [NMAX]int) {
	fmt.Printf("%2s |    %4s    |    %5s    |   %6s   | %8s | %5s | %s\n", "NO", "NAMA", "KONTAK", "LOKASI", "MATERIAL", "RATING", "RIWAYAT PELAYANAN")
	tanda := false
	for i := 0; i < arrCount; i++ {
		if s[i].nama != "" {
			fmt.Printf("%-2d | %-10s | %-12s | %-10s | %-8s |  %-.2f  | %s\n", i+1, s[i].nama, s[i].kontak, s[i].wilayah, s[i].jenismat, s[i].rating, s[i].rpelayanan[0])
			if riwayatke[i] > 0 {
				for j := 1; j <= riwayatke[i]; j++ {
					fmt.Printf("%-63s| %s\n", " ", s[i].rpelayanan[j])
				}
			}
			tanda = true
		}
	}
	if !tanda {
		fmt.Printf("\n%16sBELUM ADA DATA YANG DIINPUT !!!%17s", "", "")
	}
	fmt.Printf("\n\nADD Supplier [1] | MODIFICATION [2] | SORT [3] | SEARCH [4] | STATISTICS [5] | RETURN [6]\n")
	DaftarMitraLanjutan(s, arrCount, riwayatke)
}
func DaftarMitraLanjutan(s supplier, arrCount int, riwayatke [NMAX]int) {
	//sebagai fungsi lanjutan untuk akses opsi pada outputDaftarMitra
	var lanjutan int
	var w wilayahs
	var num int
	fmt.Scan(&lanjutan)
	switch lanjutan {
	case 1:
		inputData(&s, &arrCount)
		outputDaftarMitra(s, arrCount, riwayatke)
	case 2:
		modification(&s, riwayatke, &arrCount)
		outputDaftarMitra(s, arrCount, riwayatke)
	case 3:
		selectSorting(&s, arrCount)
		outputDaftarMitra(s, arrCount, riwayatke)
	case 4:
		selectSearch(s, arrCount, riwayatke)
		outputDaftarMitra(s, arrCount, riwayatke)
	case 5:
		countWilayah(s, &w, arrCount, &num, riwayatke)
	case 6: return
	}
}
func modification(s *supplier, riwayatke [NMAX]int, arrCount *int) {
	var replace string
	var num, num2, choice, category int
	var r3place float64
	fmt.Printf("\n MODIFICATION [1] | DELETE [2]")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("\nNAMA [1] | KONTAK [2] | JENIS MATERIAL [3] | WILAYAH [4] | RIWAYAT PELAYANAN [5] | RATING [6]\nVariabel yang ingin di modifikasi (tulis dalam ANGKA): ")
		fmt.Scan(&category)
		fmt.Scan(&num)
		if category >= 1 && category <= 5 {
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
			case 5:
				if riwayatke[num] == 0 {
					s[num].rpelayanan[0] = replace
				} else {
					fmt.Scan(&num2)
					modificationRP(s, num, num2)
				}
			}
		} else if category == 6 {
			fmt.Scan(&r3place)
			s[num].rating = r3place
		} else {
			fmt.Print("KATEGORI TIDAK DITEMUKAN !! ")
		}
	case 2:
		var idxdel int
		fmt.Printf("\nSupplier yang ingin di DELETE:")
		fmt.Scan(&idxdel)
		for i := idxdel; i < *arrCount-1; i++ {
			s[i] = s[i+1]
		}
		s[*arrCount-1] = suppliers{}
		*arrCount--
	}
}
func modificationRP(s *supplier, num int, num2 int) {
	var replace string
	fmt.Scan(&replace)
	s[num].rpelayanan[num2] = replace
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
	// selection sort ascending based on nama (untuk kebutuhan BinarySearchNama)
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
func selectSearch(s supplier, arrCount int, riwayatke [NMAX]int) {
	//fungsi untuk memilih SEARCH BY(name or lokasi...mungkin akan ditambahkan opsi lain setelah review)
	fmt.Printf("\nSEARCH BY NAME [1] | LOCATION [2] :")
	var find string
	var option int
	fmt.Scan(&option)
	fmt.Scan(&find)
	switch option {
	case 1:
		//sort by nama dulu (pakai copy s, jadi array asli tidak berubah)
		//baru binary search, binary search hanya valid pada array terurut
		SelectionSortNama(&s, arrCount)
		idx := BinarySearchNama(s, arrCount, find)
		if idx == -1 {
			fmt.Printf("Supplier '%s' tidak ditemukan.\n", find)
		} else {
			var single supplier
            var singleRiwayat [NMAX]int
			single[0] = s[idx]
			//riwayatke tidak ikut di-sort, jadi ambil dari array asli berdasarkan nama
			tanda := false
			for i := 0; i < arrCount && !tanda; i++ {
				if s[idx].nama == s[i].nama {
					singleRiwayat[0] = riwayatke[i]
					tanda = true
				}
			}
			outputDaftarMitra(single, 1, singleRiwayat)
		}
	case 2:
		SequentialSearchLokasi(s, arrCount, find, riwayatke)
	}
}
func SequentialSearchLokasi(s supplier, arrCount int, find string, riwayatke [NMAX]int) {
	//mencari supplier via lokasi menggunakan sequential search, memasukkan supplier yang memiliki lokasi yg sama ke array lokasi lalu output
	var lokasi supplier
	var riwayatLokasi [NMAX]int
	var arrLokasi int
	for i := 0; i < arrCount; i++ {
		if s[i].wilayah == find {
			lokasi[arrLokasi] = s[i]           //dilakukan untuk memindahkan isi dari supplier dicari (via lokasi) ke array lokasi pada indeks ke arrLokasi (supplier yang berbeda bisa mempunyai lokasi yang sama, maka dibuat array baru agar pada proses output bisa keluar semua supplier dengan lokasi yang dicari)
			riwayatLokasi[arrLokasi] = riwayatke[i] //memindahkan jumlah riwayat pelayanan dari supplier asli ke array riwayatLokasi agar riwayat tetap sesuai dengan supplier yang dicari
			arrLokasi++                        //menambah jumlah data yang berhasil ditemukan sekaligus menjadi indeks berikutnya untuk array lokasi
		}
	}
	outputDaftarMitra(lokasi, arrLokasi, riwayatLokasi)
}
func BinarySearchNama(s supplier, arrCount int, find string) int {
	//mencari supplier berdasarkan nama, dikarenakan nama supplier unique untuk setiap supplier maka yang keluar pasti hanya ada 1 supplier untuk setiap pencarian berdasarkan nama.
	//fungsi pencegah untuk inputan nama yang double BELUM dibuat
	//menerima array yang sudah di-sort by nama (dilakukan di selectSearch via SelectionSortNama)
	var left, mid, right int
	left = 0
	right = arrCount - 1
	for left <= right {
		mid = (left + right) / 2
		if s[mid].nama == find {
			return mid
		} else if s[mid].nama > find {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func countWilayah(s supplier, w *wilayahs, arrCount int, countw *int, riwayatke [NMAX]int) {
	//menghitung banyaknya wilayah yang ada dan masukkan ke dalam array wilayah.nama (tanpa dupe/double)
	var tanda bool
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
				w[i].riwayatpw[w[i].wilcount] = riwayatke[j]
				w[i].wilcount++
			}
		}
	}
	outputWilayah(w, *countw)
	DaftarMitraLanjutan(s, arrCount, riwayatke)
}
func outputWilayah(w *wilayahs, countw int) {
	//output jumlah dan rata2 wilayah
	for i := 0; i <= countw; i++ {
		fmt.Printf("WIlayah %s: jumlah %d supplier dan rata-rata rating adalah %.2f\n", w[i].nama, w[i].wilcount, rata2rating(w, i))
		for j := 0; j < w[i].wilcount; j++ {
			fmt.Printf("%-2d | %-10s | %-12s | %-10s | %-8s |  %-.2f  | %s\n", j+1, w[i].isiSupplier[j].nama, w[i].isiSupplier[j].kontak, w[i].isiSupplier[j].wilayah, w[i].isiSupplier[j].jenismat, w[i].isiSupplier[j].rating, w[i].isiSupplier[j].rpelayanan[0])
			if w[i].riwayatpw[j] > 0 {
				for k := 1; k <= w[i].riwayatpw[j]; k++ {
					fmt.Printf("%-63s| %s\n", " ", w[i].isiSupplier[j].rpelayanan[k])
				}
			}
		}
	}
	fmt.Printf("ADD Supplier [1] | MODIFICATION [2] | SORT [3] | SEARCH [4] | STATISTICS [5] | RETURN [6]")
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

