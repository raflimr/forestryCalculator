// Forestry Calculator adalah rumus perhitungan yang telah diprogram sehingga
// hanya memasukkan paramater nya saja, nilai dari perhitugannya akan keluar
// Untuk sementara waktu hanya bisa dijalankan dalam CLI
package forestryCalculator

import (
	"fmt"
	"strconv"
)

type JenisData struct {
	derajat, menit, detik float64
}

// Rumus Perhitungan Pencarian Decimal Degree
func DMS(derajat, menit, detik float64) float64 {
	menit = menit / 60
	detik = detik / 3600
	dms := derajat + menit + detik
	return dms
}

// Rumus Pencarian Zona UTM
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
