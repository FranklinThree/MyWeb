package main

type Elements []interface{}
type SuperString struct {
	elements Elements
	length   int
	result   string
}

func (ss *SuperString) append(e interface{}) {
	ss.elements = append(ss.elements, e)
	ss.length++
}
func (ss *SuperString) deleteFirst(e interface{}) {
	for i, element := range ss.elements {
		if element == e {
			for j := i; j < ss.length; j++ {
				ss.elements[j] = ss.elements[j+1]
			}
		}
	}
}
