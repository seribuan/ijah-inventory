package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

//"github.com/jinzhu/gorm"
//_ "github.com/jinzhu/gorm/dialects/sqlite"

func main() {
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
	//db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	//defer db.Close()

	//if err {
	//	fmt.Errorf("Something went wrong %v:", err)
	//}

}
