package universal

import "io"

//CheckErr true为正常，false为有错误
func CheckErr(err error) bool {
	if err != nil {
		//以下列出屏蔽的错误，如遇到将视为正常
		switch err {
		/*
			读取时到达最后一行，视为读取结束的标志
		*/
		case io.EOF:
			return true
		}
		panic(err)
		return false
	}
	return true
}
