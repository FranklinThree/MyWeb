package input

type InputFunction struct {
	Input
	IsFixed  []int
	IsMatch  []int
	Params   []interface{}
	function func(...interface{})
}
