package MyWeb

import (
	"strconv"
)

func Uint2String(number uint) string {
	return strconv.Itoa(int(number))
}
