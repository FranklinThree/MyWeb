package main

import (
	"database/sql/driver"
)

type Question struct {
	Id          uint
	Name        string
	Description string
	Choices     Choices
}

func (q *Question) ToStructure() (res string, err error) {
	//res = q.id + " " + q.name+"{"
	//for _, choice := range q.choices {
	//
	//
	//}
	return "", nil

}
func (q *Question) Scan(value interface{}) (err error) {
	return nil
}
func (q *Question) Value() (dv driver.Value, err error) {
	var s = ""
	for _, choice := range q.Choices.Values {
		s += Uint2String(choice.Id) + " " + choice.toString() + "\n"
	}
	return s, nil

}
