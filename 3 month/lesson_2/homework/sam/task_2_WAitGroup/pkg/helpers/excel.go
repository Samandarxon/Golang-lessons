package helpers

import (
	"app/modules"
	"fmt"
	"strconv"
	"sync"

	"github.com/xuri/excelize/v2"
)

func SaveToExcel(wg *sync.WaitGroup, inventory []modules.Money, fileName string) error {
	f := excelize.NewFile()
	defer func() {
		wg.Done()
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := fmt.Sprint(inventory[0].Date)
	fmt.Println(sheetName)
	err := f.SetSheetName("Sheet1", sheetName)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue(sheetName, "A1", "ID")
	f.SetCellValue(sheetName, "B1", "Code")
	f.SetCellValue(sheetName, "C1", "Ccy")
	f.SetCellValue(sheetName, "D1", "CcyNm_RU")
	f.SetCellValue(sheetName, "E1", "CcyNm_UZ")
	f.SetCellValue(sheetName, "F1", "CcyNm_UZC")
	f.SetCellValue(sheetName, "G1", "CcyNm_EN")
	f.SetCellValue(sheetName, "H1", "Nominal")
	f.SetCellValue(sheetName, "I1", "Rate")
	f.SetCellValue(sheetName, "J1", "Diff")
	f.SetCellValue(sheetName, "K1", "Date")
	for i := 0; i < len(inventory); i++ {
		w_g := sync.WaitGroup{}
		w_g.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// fmt.Println(inventory[i])
			f.SetCellValue(sheetName, "A"+strconv.Itoa(i+2), inventory[i].Id)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(i+2), inventory[i].Code)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(i+2), inventory[i].Ccy)
			f.SetCellValue(sheetName, "D"+strconv.Itoa(i+2), inventory[i].CcyNm_RU)
			f.SetCellValue(sheetName, "E"+strconv.Itoa(i+2), inventory[i].CcyNm_UZ)
			f.SetCellValue(sheetName, "F"+strconv.Itoa(i+2), inventory[i].CcyNm_UZC)
			f.SetCellValue(sheetName, "G"+strconv.Itoa(i+2), inventory[i].CcyNm_EN)
			f.SetCellValue(sheetName, "H"+strconv.Itoa(i+2), inventory[i].Nominal)
			f.SetCellValue(sheetName, "I"+strconv.Itoa(i+2), inventory[i].Rate)
			f.SetCellValue(sheetName, "J"+strconv.Itoa(i+2), inventory[i].Diff)
			f.SetCellValue(sheetName, "K"+strconv.Itoa(i+2), inventory[i].Date)
		}(&w_g)
		w_g.Wait()
	}
	if len(fileName) < 1 {
		fileName = "inventory"
	}
	if err := f.SaveAs(fileName + ".xlsx"); err != nil {
		fmt.Println(fileName)
		return err
	}
	return nil
}
