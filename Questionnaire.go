package main

import (
	"database/sql/driver"
	_ "gorm.io/gorm"
)

type Questionnaire struct {
	Id          uint `gorm:"primaryKey"`
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
func (qnn *Questionnaire) Scan(value interface{}) (err error) {
	return nil
}
func (qnn *Questionnaire) Value() (dv driver.Value, err error) {
	res := ""
	for _, question := range qnn.Questions {
		temp, err := question.ToStructure()
		if !CheckErr(err) {
			return nil, err
		}
		res += temp

	}
	return res, nil
}
