package questionnaire

type AnswerPaper struct {
	Id      uint `gorm:"AUTOINCREMENT"`
	Answers []Answer
}
