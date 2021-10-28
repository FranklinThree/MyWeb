package main

type Annotation struct {
	Id       uint
	Sentence string
}

func (ann *Annotation) ToStructure() (res string, err error) {
	return "", nil
}
