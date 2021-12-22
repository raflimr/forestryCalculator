// Forestry Calculator adalah rumus perhitungan yang telah diprogram jadi hanya dengan
// menginput sesuai instruksi dalam program, maka akan menghasil Zona UTM atau Decimal Degree yang sedang dicari
package forestryCalculator

import (
	"fmt"
	"strconv"
)

// Rumus Perhitungan Pencarian Decimal Degree
// Decimal Degree = Derajat + Menit + Detik
// Contoh
// Derajat = 108 Tidak perlu diubah
// Menit = Menit/60 diubah ke jam
// Detik = Detik/3600 diubak ke jam
func DMS(derajat, menit, detik float64) float64 {
	menit = menit / 60
	detik = detik / 3600
	dms := derajat + menit + detik
	return dms
}

// Rumus Pencarian Zona UTM
// Zona UTM = Decimal Degree/6 + 30(Jika memilih BT) + (nama Lintang (S or N)
// Tapi saat hasil pembagian memiliki koma nilai zona UTM ditambah 1 dan jika tidak tetap mulai dari 0
func FindUTM(derajat, menit, detik float64, lintang, bujur string) string {
	var zona float64
	var zonaInt int
	var zonaString string

	menit = menit / 60
	detik = detik / 3600
	var dcg float64 = derajat + menit + detik
	dcg = dcg / 6
	dcgString := strconv.FormatFloat(dcg, 'f', -1, 64)
	if len(dcgString) > 2 {
		zonaInt += 1
	} else {
		zonaInt = 0
	}

	switch {
	case lintang == "ls" || lintang == "LS":
		lintang = "S"
	case lintang == "lu" || lintang == "LU":
		lintang = "N"
	default:
		fmt.Println("Maaf Lintang yang anda input tidak terdaftar")
	}

	switch {
	case bujur == "bt" || bujur == "BT":
		zona = dcg + 30 + float64(zonaInt)
		zonaInt = int(zona)
		zonaString = strconv.Itoa(zonaInt)
	case bujur == "bb" || bujur == "BB":
		var zona float64 = dcg + float64(zonaInt)
		zonaInt = int(zona)
		zonaString = strconv.Itoa(zonaInt)
	default:
		fmt.Println("Maaf bujur yang anda input tidak terdaftar")
	}

	return zonaString + lintang
}
