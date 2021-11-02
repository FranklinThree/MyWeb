package main

import ()

type AnswerPaper struct {
	Id      uint `gorm:"AUTOINCREMENT"`
	Answers []Answer
}
