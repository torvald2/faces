package services

import (
	"time"
)

func SendJornalByMail(start, end time.Time, emails string, jg JornalGetter, rs ReportSender, sh SheetCreator) error {
	jornal, err := jg.GetJornalRecords(start, end)
	if err != nil {
		return err
	}
	excel, err := sh(jornal)
	if err != nil {
		return err
	}
	err = rs(excel, "Report.xlsx", emails)
	return err
}
