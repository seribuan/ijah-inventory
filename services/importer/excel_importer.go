package importer

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type ImportBarangWriter interface {
	Write(sku string, namaBarang string, jumlah int)
}

type ImportBarangMasukWriter interface {
	Write(waktu time.Time, sku string, jumlahPesan int, jumlahDiterima int, hargaBeli float64, noKwitansi string, catatan string)
}

type ImportBarangKeluarWriter interface {
	Write(waktu time.Time, sku string, jumlahKeluar int, hargaJual float64, catatan string)
}

//Import from Excel
func ImportBarang(xlsName string, barangWriter ImportBarangWriter) {
	xlFile, err := xlsx.OpenFile(xlsName)
	if err != nil {
		fmt.Errorf("Error %v", err)
	}
	//import barang part
	sheetBarang := xlFile.Sheets[0]
	rowBarang := make([]string, 3)
	for i, row := range sheetBarang.Rows {
		for j, cell := range row.Cells {
			rowBarang[j] = cell.String()
		}
		//if SKU Empty, or the title, just continue
		if rowBarang[0] == "" || i == 0 {
			continue
		}
		jumlah, _ := strconv.Atoi(rowBarang[2])
		barangWriter.Write(rowBarang[0], rowBarang[1], jumlah)

	}

}

func ImportBarangMasuk(xlsName string, barangMasukWriter ImportBarangMasukWriter) {
	//import barang masuk part
	xlFile, _ := xlsx.OpenFile(xlsName)
	sheetBarangMasuk := xlFile.Sheets[1]
	rowBarangMasuk := make([]string, 9)
	for i, row := range sheetBarangMasuk.Rows {
		for j, cell := range row.Cells {
			rowBarangMasuk[j] = cell.String()
		}
		//if SKU Empty, or the title, just continue
		if i == 0 || rowBarangMasuk[1] == "" {
			continue
		}

		waktu, _ := time.Parse("2006/01/02 15:04", rowBarangMasuk[0])
		jumlahPesan, _ := strconv.Atoi(rowBarangMasuk[3])
		jumlahDiterima, _ := strconv.Atoi(rowBarangMasuk[4])
		hargaBeli, _ := strconv.ParseFloat(rowBarangMasuk[5], 64)
		barangMasukWriter.Write(waktu, rowBarangMasuk[1], jumlahPesan, jumlahDiterima, hargaBeli, rowBarangMasuk[6], rowBarangMasuk[7])

	}

}

func ImportBarangKeluar(xlsName string, barangKeluarWriter ImportBarangKeluarWriter) {
	//import barang keluar part
	xlFile, _ := xlsx.OpenFile(xlsName)
	sheet := xlFile.Sheets[2]
	rowBarangKeluar := make([]string, 9)
	for i, row := range sheet.Rows {
		for j, cell := range row.Cells {
			rowBarangKeluar[j] = cell.String()
		}
		//if SKU Empty, or the title, just continue
		if i == 0 || rowBarangKeluar[1] == "" {
			continue
		}

		waktu, _ := time.Parse("2006-01-02 15:04:05", rowBarangKeluar[0])
		jumlahKeluar, _ := strconv.Atoi(rowBarangKeluar[3])
		hargaJual, _ := strconv.ParseFloat(rowBarangKeluar[4], 64)
		barangKeluarWriter.Write(waktu, rowBarangKeluar[1], jumlahKeluar, hargaJual, rowBarangKeluar[6])

	}
}
