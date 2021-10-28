package main

import (
	"fmt"
	"github.com/FranklinThree/MyWeb/resource"
)

func main() {
	fmt.Println("Good afternoon!")
	AServer := resource.AwesomeServer{}
	err := AServer.Start()
	resource.CheckErr(err)
}
