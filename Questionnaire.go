package main

import (
	_ "gorm.io/gorm"
)

type Questionnaire struct {
	ID          uint `gorm:"AUTOINCREMENT"`
	Info        uint
	Name        string
	Description string
	Questions   []Question
}

func (qnn *Questionnaire) ToStructure() (res string, err error) {
	//res = Uint2String(qn.id) + " " + qn.name + "{"
	//for _, obj := range qn.objects {
	//	res += "\n"
	//	temp, err := obj.ToStructure()
	//	if !CheckErr(err) {
	//		return "", err
	//	}
	//	res += temp
	//	res += "\n"
	//}
	return res, nil
}
