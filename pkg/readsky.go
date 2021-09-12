package pkg

import (
	"autosky/models"
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

func Readsky() {
	fmt.Printf("открываем файл\n")
	wb, err := xlsx.OpenFile("sky.xlsx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("открыли файл\n")
	sh, ok := wb.Sheet["Лист1"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	fmt.Printf("MaxRow=%v\n", sh.MaxRow)

	for i := 1; i < sh.MaxRow; i++ {

		theCell, err := sh.Cell(i, 7) //firm7
		if err != nil {
			continue
			panic(err)
		}
		firm, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 9) //partnumber
		if err != nil {
			continue
			panic(err)
		}
		partnumber, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 11) //name
		if err != nil {
			continue
			panic(err)
		}
		name, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 16) //price
		if err != nil {
			continue
			panic(err)
		}
		price, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 15) //qty
		if err != nil {
			continue
			panic(err)
		}
		qty, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 20) //Amount
		if err != nil {
			fmt.Printf("Amaunt err %v", theCell)
			continue
			panic(err)
		}
		amount, err := theCell.FormattedValue()
		if err != nil {
			fmt.Printf("Amaunt err %v", theCell)
			//continue
			panic(err)

		}
		theCell, err = sh.Cell(i, 3) //remark
		if err != nil {
			continue
			panic(err)
		}
		remark, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}

		tmp := models.Tiss{Firm: firm, PartNumber: partnumber, Name: name, Price: price, Qty: qty, Amount: amount, Remark: remark}
		models.Price = append(models.Price, tmp)
	}
	fmt.Printf("Price: %v\n", models.Price)
}
