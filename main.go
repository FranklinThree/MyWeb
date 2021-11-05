package main

func main() {
	as := AwesomeServer{}
	var err error
	globalConfig, err := NewConfig("global.config", "entrance")
	as.sqlConfig, err = NewConfig(globalConfig.Map["sqlConfig"], "null")
	as.netConfig, err = NewConfig(globalConfig.Map["netConfig"], "null")
	as.wsConfig, err = NewConfig(globalConfig.Map["wsConfig"], "null")
	err = as.New()
	if !CheckErr(err) {
		return
	}
	err = as.Start()
	CheckErr(err)

}
