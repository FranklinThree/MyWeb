package MyWeb

type Question struct {
	id          uint
	name        string
	description string
	choices     []Choice
}

func (q *Question) ToStructure() (res string, err error) {
	//res = q.id + " " + q.name+"{"
	//for _, choice := range q.choices {
	//
	//
	//}
	return "", nil

}
