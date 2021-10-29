package main

import "database/sql/driver"

type Choice struct {
	Id       uint
	Sentence string
}

func (c *Choice) toString() string {
	return ""
}

func (c *Choice) Scan(value interface{}) (err error) {
	return nil
}
func (c *Choice) Value() (dv driver.Value, err error) {
	return c.Sentence, nil
}

type Choices struct {
	Values []Choice
}

func (cs *Choices) Scan(value interface{}) (err error) {
	return nil

}
func (cs *Choices) Value() (dv driver.Value, err error) {
	res := ""
	for _, c := range cs.Values {
		res += c.toString()
	}
	return
}
