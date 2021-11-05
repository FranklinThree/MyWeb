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
	expr := make([]string, 1280)

	expr = append(expr, "}")
	return res, nil
}

//func Scan(expr string)(err error){
//
//}
