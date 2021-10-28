package main

func main() {
	as := AwesomeServer{}
	var err error
	as.sqlConfig, err = NewConfig("sql.config", "?")
	as.netConfig, err = NewConfig("net.config", "?")

	err = as.Start()
	CheckErr(err)

}
