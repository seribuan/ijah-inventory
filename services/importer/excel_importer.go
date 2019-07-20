package importer

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

type ImportBarangWriter interface {
	Write(SKU string, namaBarang string, jumlah int)
}

//Import from Excel
func Import(xlsName string, barangWriter ImportBarangWriter) {
	xlFile, err := xlsx.OpenFile(xlsName)
	if err != nil {
		fmt.Errorf("Error %v", err)
	}

	//for _, sheet := range xlFile.Sheets {
	sheetBarang := xlFile.Sheets[0]
	rowBarang := make([]string, 3)
	for i, row := range sheetBarang.Rows {
		for j, cell := range row.Cells {
			rowBarang[j] = cell.String()
		}
		//if SKU Empty, or the title, just continue
		if rowBarang[0] == "" || i == 1 {
			continue
		}
		jumlah, _ := strconv.Atoi(rowBarang[2])
		barangWriter.Write(rowBarang[0], rowBarang[1], jumlah)

	}
	//}
}
