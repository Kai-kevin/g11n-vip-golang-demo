package sample

import (
	"vipgoclient/com/vmware/i18n/client/util"
	"time"
	"vipgoclient/com/vmware/i18n/client/constants"
)

type Formator struct{
	Locale string
}

func (formator *Formator)GetNumber(number float64) (string,error){

	format,err := util.GetNumberFormatByLocal(formator.Locale)

	if err != nil{
		return "",err
	}

	return util.FormatNumber(&format,number),nil
}

func (formator *Formator)GetPercentNumber(number float64) (string,error){
	format,err := util.GetNumberFormatByLocal(formator.Locale)

	if err != nil{
		return "",err
	}

	return util.FormatPercent(&format,number),nil
}

func (formator *Formator)GetCurrencyNumber(number float64) (string,error){
	format,err := util.GetCurrencyFormatByLocal(formator.Locale)

	if err != nil{
		return "",err
	}

	return util.FormatCurrency(&format,number),nil
}

//@Pattern The format pattern is from cldr,currently just support patterns listed below;
//the enum value is
// SHORTDATE
//	MEDIUMDATE
//	FULLDATE
//	LONGDATE
//	SHORTTIME
//	MEDIUMTIME
//	LONGTIME
//	FULLTIME
//	SHORT
//	MEDIUM
//	LONG
//	FULL, you can use likes this: constants.SHORTDATE
func (formator *Formator) GetDate(datetime time.Time,pattern int) (string,error){
	return util.FormatDateTime(pattern,datetime,formator.Locale)
}

func (formator *Formator) GetFullDateTime(datetime time.Time) (string,error){
	return util.FormatDateTime(constants.FULL,datetime,formator.Locale)
}

func (formator *Formator) GetFullDate(datetime time.Time) (string,error){
	return util.FormatDateTime(constants.FULLDATE,datetime,formator.Locale)
}