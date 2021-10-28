package resource

type Questionnaire struct {
	id          uint
	name        string
	description string
	objects     []QuestionnaireObject
}

func (qn *Questionnaire) ToStructure() (res string, err error) {
	res = Uint2String(qn.id) + " " + qn.name + "{"
	for _, obj := range qn.objects {
		res += "\n"
		temp, err := obj.ToStructure()
		if !CheckErr(err) {
			return "", err
		}
		res += temp
		res += "\n"
	}
	return res, nil
}
