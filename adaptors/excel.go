package adaptors

import (
	"bytes"
	"fmt"
	"io"

	"atbmarket.comfaceapp/models"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateSheet(operations []models.JornalOperation) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	f := excelize.NewFile()
	index := f.NewSheet("Main")
	//Header
	f.SetCellValue("Main", "A1", "ФИО сотрудника")
	f.SetCellValue("Main", "B1", "Номер магазина")
	f.SetCellValue("Main", "C1", "Дата")
	f.SetCellValue("Main", "D1", "Время")
	f.SetCellValue("Main", "E1", "Приход/Уход")
	for i, op := range operations {
		f.SetCellValue("Main", fmt.Sprintf("A%v", i+2), op.UserName)
		f.SetCellValue("Main", fmt.Sprintf("B%v", i+2), op.ShopNum)
		f.SetCellValue("Main", fmt.Sprintf("C%v", i+2), op.OperationDate.Format("2006.01.02"))
		f.SetCellValue("Main", fmt.Sprintf("D%v", i+2), op.OperationDate.Format("03:04:05"))
		if op.OperationType == models.Coming {
			f.SetCellValue("Main", fmt.Sprintf("E%v", i+2), "Приход")
		} else {
			f.SetCellValue("Main", fmt.Sprintf("E%v", i+2), "Уход")
		}

	}
	f.SetActiveSheet(index)
	err := f.Write(buf)
	return buf, err

}
