package main

type Question struct {
	Id          uint
	Name        string
	Description string
	Choices     []Choice
}

func (q *Question) ToStructure() (res string, err error) {
	//res = q.id + " " + q.name+"{"
	//for _, choice := range q.choices {
	//
	//
	//}
	return "", nil

}
