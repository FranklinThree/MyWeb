package MyWeb

type QuestionnaireObject interface {
	ToStructure() (string, error)
}
