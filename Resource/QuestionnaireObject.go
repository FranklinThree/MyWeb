package resource

type QuestionnaireObject interface {
	ToStructure() (string, error)
}
