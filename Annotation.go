package MyWeb

type Annotation struct {
	id       uint
	sentence string
}

func (ann *Annotation) ToStructure() (res string, err error) {
	return "", nil
}
