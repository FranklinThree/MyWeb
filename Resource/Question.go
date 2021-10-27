package Resource

type Question struct {
	id      uint
	ask     string
	choices []Choice
}

func (q *Question) ToStructure() {

}
