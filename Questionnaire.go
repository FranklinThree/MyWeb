package main

import (
	_ "gorm.io/gorm"
)

type Questionnaire struct {
	Id          uint                  `gorm:"column:id primaryKey"`
	Name        string                `gorm:"column:name"`
	Description string                `gorm:"column:description"`
	Objects     []QuestionnaireObject `gorm:"column:objects"`
}

func (qn *Questionnaire) ToStructure() (res string, err error) {
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
