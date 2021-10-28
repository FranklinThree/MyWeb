package resource

type Question struct {
	id          uint
	name        string
	description string
	choices     []Choice
}

func (q *Question) ToStructure() (res string, err error) {
	res = ""
	for _, choice := range q.choices {

	}

}
