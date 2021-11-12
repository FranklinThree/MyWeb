package questionnaire

import (
	"MyWeb/universal"
	_ "gorm.io/gorm"
	"os"
)

type Questionnaire struct {
	ID          uint `gorm:"AUTOINCREMENT"`
	Info        uint
	Name        string
	Description string
	Questions   []Question
}

func (qnn *Questionnaire) ToStructure() (res string, err error) {
	expr := make([]string, 1280)

	expr = append(expr, "}")
	return res, nil
}

//func Scan(expr string)(err error){
//
//}

const (
	File = 1
	Text = 2
)

func Scan(Type int, value string) (qnn *Questionnaire, err error) {
	if Type == File {
		var reader *os.File
		reader, err = os.Open(value)
		if !universal.CheckErr(err) {
			universal.ConsolePrint(universal.Error, "Failed to read file", "fileName", value, "error", err)
			return
		}
		buffer := make([]byte, 1024*1024)
		value = ""
		for {
			var readcount int
			readcount, err = reader.Read(buffer)
			if !universal.CheckErr(err) {
				break
			}
			value += string(buffer[:readcount])
		}
	}
	for {

	}

}
