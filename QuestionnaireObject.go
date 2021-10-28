package main

type QuestionnaireObject interface {
	ToStructure() (string, error)
}
