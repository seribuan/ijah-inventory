package services

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

//Import from Excel
func Import(xlsName string) {
	xlFile, err := xlsx.OpenFile("Toko Ijah.xlsx")
	if err != nil {
		fmt.Errorf("Error %v", err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
