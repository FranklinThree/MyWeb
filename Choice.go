package main

type Choice struct {
	ID         uint `gorm:"AUTOINCREMENT"`
	Info       uint
	Sentence   string
	QuestionID uint
}

func (c *Choice) toString() string {
	return ""
}
