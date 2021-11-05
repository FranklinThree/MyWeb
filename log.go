package main

import "fmt"

const (
	Info    = 0
	Warning = 1
	Error   = 2
)

func ConsolePrint(Type int, a ...interface{}) {
	var ConsolePrintType = []string{
		"INFO", "WARNING", "ERROR",
	}
	fmt.Println("["+ConsolePrintType[Type]+"]", a)
}

//func PrintInfo(a ...interface{}){
//	fmt.Println("[INFO]",a)
//}
//func PrintWarning(a ...interface{}){
//	fmt.Println("[WARNING]",a)
//}
//func PrintError(a ...interface{}){
//	fmt.Println("[ERROR]",a)
//}
