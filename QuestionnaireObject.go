package main

import "database/sql/driver"

type QuestionnaireObject interface {
	ToStructure() (string, error)
	Scan(value interface{}) error
	Value() (driver.Value, error)
}
