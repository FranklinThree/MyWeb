package questionnaire

type Question struct {
	ID              uint `gorm:"AUTOINCREMENT"`
	Info            uint
	Name            string
	Description     string
	Choices         []Choice
	QuestionnaireID uint
}

func (q *Question) ToStructure() (res string, err error) {
	//res = q.id + " " + q.name+"{"
	//for _, choice := range q.choices {
	//
	//
	//}
	return "", nil
}
